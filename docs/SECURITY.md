# Security Considerations

Graveyard, a local terminal application, manages API keys for external services. The security of these keys depends on the underlying operating system and filesystem permissions, not on encryption by the application itself.

## How Graveyard Stores API Keys

### API Key Storage Implementation

1.  **Location**: API keys are stored in a local `.env` file in the project's root directory.
2.  **Format**: The file contains plaintext environment variables (e.g., `GEMINI_API_KEY=your_key_here`).
3.  **Permissions**: The file is created with strict filesystem permissions (`0600` on Unix-like systems, which means read/write for the owner only). This is set programmatically by the `ConfigService` in `internal/service/config.go`.
4.  **Lifecycle**: Keys are loaded at startup and can be added, updated, or deleted through the application's Settings UI or by manually editing the `.env` file.

### Why Plaintext Storage?

This is an intentional design choice based on the threat model and practical considerations for a local CLI tool:

1.  **Industry Standard**: Most local CLI tools (Docker, AWS CLI, kubectl) store credentials in plaintext configuration files.
2.  **User Control**: Users have full control over their configuration files, can easily edit them, and are responsible for their security.
3.  **Limited Threat Surface**: The primary security concern is physical access to the machine or unauthorized access to the user's account. In these scenarios, plaintext storage in a secure file does not significantly increase the risk.
4.  **False Security**: Implementing encryption within the application would be ineffective, as the encryption key would have to be stored alongside or within the executable, rendering the encryption pointless.
5.  **No Inter-Process Communication**: The application does not share the keys with any other process, keeping them isolated to the Graveyard application and the user who owns the `.env` file.

## Security Measures in Place

Graveyard implements several measures to protect API keys:

1.  **Strict File Permissions**: The `.env` file is created with `0600` permissions (owner read/write only), preventing other users on a multi-user system from reading the file.
2.  **`.gitignore` Protection**: The `.env` file is included in `.gitignore` by default, preventing accidental commits to version control.
3.  **No Logging of Secrets**: The application never logs the full API keys to the console or to `graveyard.log`. The Settings UI may show a truncated version (e.g., "AIzaSyD...") for user verification.
4.  **Secure Key Deletion**: The "Delete All" function in Settings securely removes the entire `.env` file.

## Best Practices for Users

As a user of Graveyard, you play a crucial role in the security of your API keys:

1.  **Secure Your Machine**: Ensure your computer is protected against malware and that you have proper physical security.
2.  **Protect the `.env` File**: Never share the `.env` file or paste its contents in public forums, support requests, or screenshots.
3.  **Version Control**: Do not commit the `.env` file to Git or any other version control system.
4.  **Use Separate Keys**: Consider creating and using separate API keys for different machines or environments to limit the impact of a potential compromise.
5.  **Monitor API Usage**: Regularly check your Google AI Studio and VirusTotal dashboards for any unexpected or suspicious activity.
6.  **Rotate Keys**: Periodically generate new API keys and update your Graveyard configuration, especially if you suspect your key has been compromised or if you have shared your screen or granted remote access to someone else.

## What Graveyard Protects Against

Graveyard's security model is designed to protect against common, local threats:

1.  **Casual Snooping**: Other users on the same system cannot read your `.env` file due to the `0600` file permissions.
2.  **Accidental Exposure**: The `.gitignore` file prevents the `.env` file from being accidentally committed to version control.
3.  **Log File Leakage**: API keys are never exposed in log files or error messages.

## Limitations & What We Don't Protect Against

It's important to understand the limitations of the security model:

1.  **Malware with User Privileges**: If malicious software runs with your user account, it can read the `.env` file.
2.  **System Administrators**: Users with administrative privileges (root on Linux/macOS, Administrator on Windows) can read any file on the system, including the `.env` file.
3.  **In-Memory Access**: While the application is running, the API keys are in memory and could potentially be extracted from a memory dump of the process.
4.  **Physical Access**: If an attacker gains physical access to your computer while it's powered on or if the disk is not encrypted, they could potentially access your files.
5.  **Reverse Engineering**: A determined and skilled attacker could reverse-engineer the application, although this is a complex task for most users.

## API Provider Security

Both Google (Gemini) and VirusTotal use secure, encrypted connections (HTTPS) to transmit data, including your API keys and analysis requests. You should always ensure you are accessing the correct URLs when configuring your keys and be wary of phishing attempts.

## Recommendations for Enhanced Security

For users with higher security requirements:

1.  **Use System Keychains**: Consider using your operating system's built-in keychain (macOS Keychain, Windows Credential Manager, Linux Secret Service) to store API keys. You can then set environment variables that reference these keychains before running Graveyard.
2.  **Dedicated Secrets Management**: For teams or enterprises, integrate Graveyard with a dedicated secrets management solution like HashiCorp Vault, AWS Secrets Manager, or 1Password.
3.  **Hardware Security Modules (HSMs)**: For organizations with the highest security requirements, HSMs can be used for cryptographic key storage and management.
4.  **Run in a Virtual Machine or Container**: Running Graveyard in an isolated environment (VM or container) can limit the potential impact of a compromise.
5.  **Full Disk Encryption**: Ensure your entire hard drive is encrypted (e.g., using FileVault on macOS, BitLocker on Windows, or LUKS on Linux). This protects your data at rest.

## Reporting Security Vulnerabilities

If you discover a security vulnerability in Graveyard itself:

1.  **Do not open a public issue** on GitHub.
2.  **Contact the maintainer** via a private channel or email with a detailed description of the vulnerability.
3.  **Provide time for a fix** before any public disclosure, following responsible disclosure practices.

## Compliance Notes

-   **Privacy**: Graveyard does not collect or transmit any telemetry or personal data. All API calls are made directly from your machine to the respective providers (Google, VirusTotal).
-   **Data Sent to Third Parties**: 
    -   **Gemini API**: Process details and your questions are sent to Google for analysis.
    -   **VirusTotal**: Only the SHA256 hash of the process executable is sent to VirusTotal; the file itself is never uploaded.

## Conclusion

Graveyard's security model is appropriate and robust for a local terminal application used by individual users and small teams. It prioritizes simplicity and user control while implementing standard security practices for protecting secrets on a local machine.

For users in high-security environments, the best approach is to use Graveyard as one component in a comprehensive security strategy that includes strong personal cybersecurity hygiene, OS-level security features, and potentially dedicated secrets management systems.

---

**Last Updated**: November 2025
