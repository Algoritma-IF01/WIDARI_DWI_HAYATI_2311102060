![ss perulangan1](https://github.com/user-attachments/assets/a8e3b331-f9de-4a30-83a0-d5a618e671e5)

Kode ini adalah program yang mengecek urutan warna yang benar dari pengguna. Awalnya, warna yang benar didefinisikan sebagai merah, kuning, hijau, dan ungu. Program menetapkan variabel hasil sebagai true untuk melacak apakah input pengguna benar selama lima percobaan.
Setiap percobaan meminta input warna dari pengguna empat kali, menanyakan warna pertama hingga keempat. Setelah menerima input, program membandingkan input warna pengguna dengan urutan yang benar. Jika ada warna yang salah, variabel hasil diatur menjadi false.
Setelah lima percobaan selesai, program mencetak "BERHASIL" diikuti oleh nilai dari hasil, menunjukkan apakah pengguna berhasil memberikan urutan warna yang benar untuk semua percobaan.

![ss perulangan2](https://github.com/user-attachments/assets/5bca3e1b-4e62-4bb4-ac02-5318c1dd711c)

Program ini meminta pengguna untuk memasukkan nama bunga secara bertahap. Program akan terus meminta nama bunga sampai pengguna mengetik "selesai". Setiap nama bunga yang dimasukkan akan disimpan dalam sebuah daftar yang disebut pita dan dipisahkan dengan tanda " - ". Program juga menghitung berapa banyak nama bunga yang telah dimasukkan. Setelah selesai, program mencetak daftar nama bunga dan jumlah total bunga yang telah dimasukkan.

![ss perulangan3](https://github.com/user-attachments/assets/68405ec9-c35a-4782-96dc-018fde8f4455)

Program ini dirancang untuk membantu Pak Andi memastikan berat belanjaan di kedua kantong sepeda motornya seimbang. Program akan terus meminta input dari pengguna mengenai berat belanjaan di kantong kiri dan kanan. Proses ini akan berlanjut sampai salah satu dari dua kondisi terjadi: berat salah satu kantong kurang dari 0 atau total berat kedua kantong melebihi 150 kg, yang akan mengakhiri proses dengan mencetak "Proses selesai."
Dalam setiap iterasi, program menghitung total berat belanjaan dan selisih berat antara kantong kiri dan kanan. Jika selisih beratnya mencapai 9 kg atau lebih, program menandai bahwa sepeda motor Pak Andi akan oleng (tidak seimbang) dan mencetak hasil tersebut.

![ss perulangan4](https://github.com/user-attachments/assets/6a61762a-5a94-4e77-8675-fd5c168351c0)

Program ini menghitung nilai perkiraan dari akar kuadrat 2. Pengguna diminta memasukkan sebuah bilangan bulat positif K. Nilai 2 dihitung melalui sebuah fungsi rekursif yang melakukan iterasi sebanyak K kali.
Pada setiap iterasi, program menghitung nilai numerik dari (4k+2)^2 dan nilai pembaginya dari hasil perkalian (4k+1) dan (4k+3). Hasil pembagian tersebut dikalikan dengan hasil dari iterasi sebelumnya untuk mendapatkan hasil akhir. Setelah perhitungan selesai, program akan mencetak nilai perkiraan 2 dengan ketelitian hingga 10 angka di belakang koma

![ss percabangan1](https://github.com/user-attachments/assets/cb411179-8929-4834-be30-e01b2d5e73e9)

Program ini dirancang untuk menghitung biaya pengiriman berdasarkan berat parsel yang dimasukkan oleh pengguna. Pertama, program meminta pengguna untuk memasukkan berat parsel dalam gram. Setelah menerima input, program menghitung total berat dalam kilogram dan sisa berat dalam gram.
Biaya pengiriman dihitung berdasarkan total berat dalam kilogram dengan tarif Rp. 10.000 per kg. Jika sisa berat kurang dari 500 gram, tambahan biaya dikenakan sebesar Rp. 15 per gram. Namun, jika sisa berat 500 gram atau lebih, tambahan biaya dikenakan sebesar Rp. 5 per gram.
Jika total berat melebihi 10 kg, maka sisa berat tidak dikenakan biaya tambahan. Program kemudian mencetak detail berat dalam kilogram dan gram, detail biaya pengiriman, serta total biaya pengiriman

![Screenshot 2024-10-15 160553](https://github.com/user-attachments/assets/337ec0a9-84c2-47e6-896f-02aa8082d801)

a.	Jika nam diberikan adalah 80.1, program akan error karena tipe data yang tidak sesuai yaitu "nam" yang diinisiasikan sebagai float64 tidak bisa diubah menjadi string sehingga tidak bisa mengoutputkan "A".

b.	Berikut kesalahan dari program tersebut.
-	Tipe Data Tidak Sesuai: 
nam dideklarasikan sebagai float64, tetapi digunakan untuk menyimpan nilai huruf seperti "A", "B", dll. Ini menyebabkan error karena tipe data tidak sesuai.
-	Logika Kondisional yang Tidak Konsisten:
Blok if tidak menggunakan else if, sehingga setiap kondisi dievaluasi secara terpisah, bukan berurutan. Ini menyebabkan tumpang tindih kondisi yang salah dalam menentukan nilai huruf.
-	Variabel nmk Tidak Diubah:
Variabel nmk untuk menyimpan nilai huruf tidak pernah diubah dalam blok kondisi.

alur program seharusnya.
-	Deklarasi Variabel:
Deklarasikan nam sebagai float64 untuk nilai angka dan nmk sebagai string untuk nilai huruf.
-	Input Nilai:
Minta input dari pengguna untuk nam.
-	Logika Kondisional dengan else if:
Gunakan logika else if untuk memastikan setiap kondisi dievaluasi secara berurutan dan hanya satu kondisi yang benar akan dieksekusi.
-	Cetak Hasil.
Setelah menentukan nilai huruf (nmk), cetak hasilnya.

c.	Program yang telah diperbaiki
![ss percabangan2](https://github.com/user-attachments/assets/b8951bcb-0b4f-40c8-9b52-fb39fa35713a)


![ss percabangan3](https://github.com/user-attachments/assets/3a62d666-20d1-405a-83ce-334d9988b14f)

Program ini dibuat untuk membantu pengguna menghitung faktor-faktor bilangan bulat positif dan menentukan apakah bilangan tersebut adalah bilangan prima. Proses dimulai dengan meminta input bilangan bulat dari pengguna. Jika bilangan yang dimasukkan lebih kecil atau sama dengan nol, program akan menampilkan pesan kesalahan dan berhenti.
Setelah menerima input, program menghitung dan menampilkan semua faktor dari bilangan tersebut dengan iterasi dari 1 hingga bilangan itu sendiri. Setiap faktor yang ditemukan akan dicetak ke layar. Selain itu, program juga menghitung jumlah faktor yang ditemukan.
Pada langkah terakhir, program menentukan apakah bilangan tersebut adalah bilangan prima. Jika jumlah faktor yang ditemukan adalah dua (artinya hanya 1 dan bilangan itu sendiri), maka bilangan tersebut adalah bilangan prima dan program menampilkan "Prima: true". Jika tidak, program menampilkan "Prima: false".


