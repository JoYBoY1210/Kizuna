```markdown
# 📡 Kizuna - P2P File Sharing CLI Tool

Kizuna is a simple peer-to-peer file sharing tool written in Go, inspired by BitTorrent. It allows you to **seed** and **download** large files using chunked transfers between peers.

---

## 📦 Included Files

```

kizuna/
├── kizuna.exe         # Windows executable
├── kizuna-linux       # Linux executable
├── kizuna-mac         # macOS executable
├── sample.meta        # Example metadata file (optional)
└── README.md          # You're reading this!

````

---

## 🛠 Requirements

- No dependencies needed.
- Just double-click or run the appropriate binary from the terminal.

> ✅ No Go installation required. Binaries are precompiled for each platform.

---

## 🚀 Usage

### ✅ 1. Seeding a New File

Start seeding a file on a specific port and optionally define peers:

```bash
./kizuna.exe seed --file="path/to/file.pdf" --port=6253 --peers=http://localhost:6254
````

* `--file`: Path to the file to share
* `--port`: Port to host the chunk server
* `--peers`: (Optional) Initial peer list for redundancy

A `.meta` file will be generated automatically.

---

### ✅ 2. Seeding from an Existing `.meta` File

If you already have a `.meta` file, you can resume seeding:

```bash
./kizuna.exe seed --meta="path/to/file.meta" --port=6253
```

---

### ✅ 3. Downloading a File

Use the `.meta` file to download the file from available peers:

```bash
./kizuna.exe download --meta="path/to/file.meta" --output="desired_filename.pdf"
```

* Downloads chunks in parallel from peers listed in the `.meta`.

---

## 🌐 Example Workflow

1. **Seeder (User A):**

   ```bash
   ./kizuna.exe seed --file="movie.mp4" --port=6253
   ```

2. **Send the `.meta` file to your friend (User B).**

3. **Downloader (User B):**

   ```bash
   ./kizuna.exe download --meta="movie.meta" --output="movie.mp4"
   ```

4. **User B can now also help seed by running:**

   ```bash
   ./kizuna.exe seed --meta="movie.meta" --port=6254
   ```

---

## 🖥 Platform Notes

* 🪟 `kizuna.exe` for **Windows**
* 🐧 `kizuna-linux` for **Linux**
* 🍎 `kizuna-mac` for **macOS**

Make sure the file is executable (`chmod +x kizuna-*` on UNIX systems).

---

## 💬 Contact

Built by \[Your Name].
Feel free to open issues or contribute on GitHub (if applicable).

---

## 🔐 Disclaimer

This is a learning project inspired by BitTorrent. Do **not** use it to share copyrighted or illegal material.

```