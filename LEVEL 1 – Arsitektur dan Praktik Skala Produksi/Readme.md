# Level 1 – Arsitektur dan Praktik Skala Produksi

## Review Level 1
Target: Menulis kode maintainable, modular, dan testable

Materi	Tujuan
Project structure (cmd/pkg/internal)	Struktur profesional
Dependency injection sederhana	Clean code, testability
Error wrapping & custom error types	Logging & debug jelas
Logger (logrus/zap) & config (Viper/env)	Real-world readiness
Unit test + coverage	Kode yang reliable
Benchmark test (testing.B)	Efisiensi dan optimalisasi

🧪 Latihan:

Refactor task manager menjadi cmd/cli, internal/task, pkg/common

Tambah config .env, logger, unit test