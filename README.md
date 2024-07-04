# Tugas Akhir - Sanbercode Golang 

## Dokumentasi Aplikasi
Aplikasi ini dirancang untuk membantu pemain Genshin Impact dalam menyusun tim yang optimal berdasarkan elemen, senjata, dan karakter yang tersedia. Aplikasi ini menyediakan database yang terdiri dari elemen, senjata, karakter, artefak, dan tim yang dapat dikustomisasi oleh pengguna.

### Judul: Rekomendasi Tim Genshin Impact

### Dokumentasi API: https://documenter.getpostman.com/view/813958/2sA3dxFCoa

#### Elements [icon: zap, color: yellow]
- `id`: Integer (Primary Key)
- `name`: String (Nama elemen)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

#### Weapons [icon: crosshair, color: red]
- `id`: Integer (Primary Key)
- `name`: String (Nama senjata)
- `type`: String (Tipe senjata)
- `rarity`: Integer (Kelangkaan)
- `base_attack`: Integer (Serangan dasar)
- `secondary_stat`: String (Statistik sekunder)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

#### Characters [icon: user, color: green]
- `id`: Integer (Primary Key)
- `name`: String (Nama karakter)
- `element_id`: Integer (Foreign Key ke Elements)
- `weapon_id`: Integer (Foreign  Key ke Weapons)
- `rarity`: Integer (Kelangkaan)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

#### Artifacts [icon: box, color: purple]
- `id`: Integer (Primary Key)
- `name`: String (Nama artefak)
- `description`: String (Deskripsi)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

#### Teams [icon: users, color: blue]
- `id`: Integer (Primary Key)
- `name`: String (Nama tim)
- `description`: String (Deskripsi tim)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `verify_status`: String (Status verifikasi)
- `verify_by`: String (Verifikasi oleh)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

#### Team Characters [icon: link, color: orange]
- `id`: Integer (Primary Key)
- `team_id`: Integer (Foreign Key ke Teams)
- `character_id`: Integer (Foreign Key ke Characters)
- `artifact_id`: Integer (Foreign Key ke Artifacts)
- `type_set`: String (Jenis set)
- `mechanism`: String (Mekanisme)
- `created_at`: DateTime (Waktu pembuatan)
- `updated_at`: DateTime (Waktu pembaruan)
- `created_by`: String (Pembuat)
- `updated_by`: String (Pembaruan oleh)

### Mendefinisikan hubungan antar tabel
characters.element_id -> elements.id
characters.weapon_id -> weapons.id
team_characters.team_id -> teams.id
team_characters.character_id -> characters.id
team_characters.artifact_id -> artifacts.id
