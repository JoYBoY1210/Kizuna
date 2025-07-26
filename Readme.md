# Kizuna - Peer-to-Peer File Sharing in Go

Kizuna is a peer-to-peer file sharing tool written in Go. It allows users to split files into chunks, generate metadata, seed files via HTTP, and download them from multiple peers in parallel with hash verification.

---

## 🔧 Features

- 📦 Chunk-based file sharing
- 🔐 SHA256 chunk verification
- 🌐 HTTP server for seeding
- ⚡ Parallel download from multiple peers
- 📁 Metadata file (`.meta`) with file and chunk info

---

## 📁 Folder Structure (Inside ZIP)

```
Kizuna/
├── kizuna.exe             # The executable (for Windows)
├── README.md              # This file
├── files/
│   ├── chunker.go
│   ├── hasher.go
│   └── meta.go
├── server/
│   └── server.go
├── downloads/
│   └── download.go
├── types/
│   └── meta.go
├── seed.go
├── download.go
└── main.go
```

---

## 🚀 Usage

### ✅ Seed a File

```bash
./kizuna.exe seed --file="C:\path\to\your\file.pdf" --port=6253 --peers=http://localhost:6253
```

Or use an existing `.meta` file:

```bash
./kizuna.exe seed --meta="file.pdf.meta" --port=6253
```

---

### ✅ Download a File

```bash
./kizuna.exe download --meta="file.pdf.meta" --output="output.pdf"
```

Downloads file from peers listed in `.meta` file in parallel using chunk requests.

---

## 📌 Notes

- Chunk size is 1MB.
- Make sure all seeders are running when downloading.
- Each peer must host the `chunks/` directory with chunk files.

---

## 📤 Sharing

To share the project:

1. Zip the full folder (including `.exe`, `README.md`, and all subdirectories).
2. Share via Google Drive, GitHub, or email.

---

## 📞 Credits

Built with ❤️ in Go by Tanish Mirajkar.