# Panduan Integrasi VirusTotal

## Ringkasan

Graveyard berintegrasi dengan VirusTotal untuk menyediakan alur kerja analisis keamanan yang mulus. Ketika Anda memindai proses, Graveyard tidak hanya mengkueri VirusTotal untuk hasil deteksi malware tetapi juga menggunakan Google Gemini AI untuk menafsirkan hasil dan memberikan penilaian yang jelas dan dapat ditindaklanjuti.

## Cara Kerja

Integrasi ini menggabungkan deteksi malware otomatis dengan analisis berbasis AI:

1.  **Pemindaian VirusTotal**: Graveyard menghitung hash SHA256 dari executable proses dan mengkueri database VirusTotal API v3.
2.  **Interpretasi AI**: Model Google Gemini 2.5 Flash menganalisis data VirusTotal dan memberikan rekomendasi keamanan.

Pendekatan ganda ini memberikan Anda data mentah dan interpretasi tingkat ahli, memudahkan pemahaman hasil pemindaian yang kompleks.

## Fitur Utama

-   **Deteksi Multi-Engine**: Memanfaatkan lebih dari 70 mesin antivirus melalui satu kueri VirusTotal.
-   **Kategorisasi Ancaman**: Secara otomatis mengklasifikasikan hasil menjadi Malicious, Suspicious, Harmless, atau Undetected.
-   **Analisis Berbasis Risiko**: AI menyediakan penilaian tingkat risiko dan rekomendasi tindakan spesifik.
-   **Alur Kerja Terpadu**: Hasil dari VirusTotal dan AI ditampilkan secara bersamaan di UI untuk tampilan komprehensif.

## Persiapan

### 1. Dapatkan API Key VirusTotal

1.  Kunjungi [VirusTotal.com](https://www.virustotal.com/) dan buat akun gratis.
2.  Buka [bagian API key](https://www.virustotal.com/gui/my-apikey) untuk menemukan API key privat Anda.

### 2. Dapatkan API Key Gemini

1.  Kunjungi [Google AI Studio](https://aistudio.google.com/).
2.  Klik "Get API Key" dan buat kunci di proyek baru atau yang sudah ada.

### 3. Konfigurasi Graveyard

Anda dapat mengonfigurasi API key dengan dua cara:

**Opsi A: UI Pengaturan (Direkomendasikan)**
1.  Jalankan Graveyard.
2.  Tekan `s` untuk membuka dialog Pengaturan.
3.  Tempelkan kunci Anda ke kolom yang sesuai.
4.  Klik "Simpan".

**Opsi B: Berkas `.env`**
1.  Di direktori proyek, buat berkas `.env` (atau edit jika sudah ada) dari template (`.env.example`).
2.  Tambahkan kunci Anda:
    ```
    GEMINI_API_KEY=kunci_api_gemini_anda
    VIRUSTOTAL_API_KEY=kunci_api_virustotal_anda
    ```

## Memindai Proses

1.  **Pilih Proses**: Navigasi ke proses yang ingin Anda analisis menggunakan tombol panah di Daftar Proses.
2.  **Mulai Pemindaian**: Tekan `t`.
3.  **Lihat Hasil**:
    -   **Panel Tengah (Data VirusTotal)**: Menunjukkan data pemindaian mentah, termasuk statistik deteksi, tingkat ancaman, dan deteksi antivirus teratas.
    -   **Panel Kiri (Analisis AI)**: Setelah pemindaian selesai, AI akan memberikan interpretasi, penilaian risiko, dan tindakan yang direkomendasikan.

## Memahami Hasil

### Panel VirusTotal (Tengah)

Panel ini menyajikan data mentah dari VirusTotal.

-   **Tingkat Ancaman & Ikon**: Indikator visual dari ancaman keseluruhan.
    -   ⚠ **MERAH**: Ancaman tinggi (beberapa deteksi malicious).
    -   ⚡ **KUNING**: Ancaman sedang (deteksi suspicious).
    -   ✓ **HIJAU**: Aman (tidak ada ancaman yang terdeteksi).
    -   ? **PUTIH**: Tidak diketahui (data tidak mencukupi).

-   **Informasi Proses**: Nama dan PID proses yang dipindai.
-   **Detail File**: Jalur file lengkap dan hash yang dipotong dari executable.
-   **Hasil Deteksi**: Barra progresso dan jumlah yang menunjukkan berapa dari 70+ mesin yang menandai file sebagai malicious atau suspicious.
-   **Ringkasan**: Ringkasan teks singkat dari temuan.
-   **Deteksi Teratas**: Daftar hingga 5 deteksi paling menonjol, termasuk nama engine dan nama ancaman spesifik yang dilaporkan.

### Panel AI Assistant (Kiri)

Panel ini, didukung oleh Gemini, memberikan analisis ahli.

AI akan menyediakan:
1.  **Penilaian Risiko**: Evaluasi keseluruhan tingkat ancaman.
2.  **Analisis Deteksi**: Penjelasan tentang arti deteksi, termasuk diskusi tentang potensi false positive.
3.  **Tindakan yang Direkomendasikan**: Panduan jelas tentang apa yang harus Anda lakukan selanjutnya (mis.,终止 proses, karantina file, atau lanjutkan pemantauan).
4.  **Konteks Tambahan**: Informasi tentang jenis proses dan signifikansinya pada sistem.

## Contoh Alur Kerja

### Skenario: Proses Sistem yang Mencurigakan

1.  Anda melihat `svchost.exe` (proses sistem Windows) menggunakan CPU lebih dari biasanya.
2.  Anda memilihnya dan menekan `t` untuk memindai.
3.  **Panel VirusTotal Menunjukkan**:
    -   **Tingkat Ancaman**: ⚡ (Sedang)
    -   **Deteksi**: "Engines: 7 / 70 detected"
    -   **Deteksi Teratas**: "Kaspersky: Trojan.Win32.Generic"
4.  **Panel AI Menunjukkan**:
    -   **Penilaian Risiko**: "Sedang-Rendah (kemungkinan false positive)"
    -   **Analisis**: "Svchost.exe adalah proses sistem Windows yang penting. Deteksi kemungkinan adalah false positive karena sifatnya yang kompleks, multi-layanan, dan aktivitas jaringan. Banyak mesin antivirus menandai file sistem yang sah ketika mereka menampilkan perilaku tertentu."
    -   **Rekomendasi**: "Jangan终止. Ini adalah proses sistem. Sebagai gantinya, gunakan Task Manager Windows untuk mengidentifikasi dan终止 layanan spesifik (svchost.exe -k ...) yang menyebabkan penggunaan CPU tinggi."

Ini mendemonstrasikan bagaimana AI membantu Anda menghindari kesalahan sistem yang kritis dengan mengontekstualisasi data.

## Best Practices

-   **Pindai Proses Tidak Dikenal**: Selalu pindai proses yang tidak Anda kenali sebelum mengambil tindakan.
-   **Percayai, Tapi Verifikasi**: Perlakukan deteksi VirusTotal sebagai titik awal, bukanvonis. Gunakan analisis AI untuk memahami konteks.
-   **Pertimbangkan Sumber**: Proses sistem dikenal dengan false positive. Berhati-hatilah ekstra saat berurusan dengan `svchost.exe`, `lsass.exe`, `csrss.exe`, dll.
-   **Pantau Tren**: Jika proses yang sebelumnya aman mulai menunjukkan deteksi mencurigakan, mungkin proses tersebut telah disusupi.
-   **Perhatikan Batas Tarif**: API key VirusTotal gratis terbatas pada 4 permintaan per menit.

## Batas Tarif API

-   **Tingkat Gratis**: 4 permintaan/menit, 500/hari, 15,500/bulan.
-   **Tingkat Premium**: Batas lebih tinggi dan fitur tambahan.

## Pemecahan Masalah

-   **"File tidak ditemukan di database VirusTotal"**: Ini umum untuk aplikasi kustom atau yang baru dikembangkan. File belum dipindai sebelumnya.
-   **"Batas tarif terlampaui"**: Tunggu sejenak sebelum memindai lagi. Lebih selektif tentang apa yang Anda pindai.
-   **"Layanan AI tidak tersedia"**: Pastikan API key Gemini Anda dikonfigurasi dengan benar di Pengaturan.

## Privasi

-   **Tidak Ada Upload File**: Graveyard hanya mengirim hash SHA256 ke VirusTotal; file aktual tidak pernah diunggah.
-   **Pembagian Hash**: VirusTotal dapat membagikan hash dengan komunitas keamanan yang lebih luas untuk meningkatkan kemampuan deteksi.

## Lihat Juga

-   [Ikhtisar Fitur](FEATURES.md) - Daftar fitur lengkap
-   [Tata Letak UI](UI_LAYOUT.md) - Memahami antarmuka pengguna
-   [Pertimbangan Keamanan](SECURITY.md) - Bagaimana kami melindungi API key Anda
