Challenge Webinar Golang 2 by Azka Lailatul Hana

Challenge ini merupakan challenge webinar golang 2 yang diadakan oleh Spacevolver yang dilaksanakan pada hari Minggu tanggal 15 September 2024, pada challenge ini diminta untuk:
- Membuat Struct Person dengan field name string dan age int.
- Fungsi CreatePerson yang dapat menerima nama dan umur sebagai parameter dan mengembalikan pointer ke struct Person, jika umur kurang dari 0, fungsi harus memanggil panic dengan pesan "Invalid Age".
- Fungsi PrintPerson yang menerima pointer ke struct Person dan mencetak informasi person.
- Fungsi HandlePerson yang memanggil CreatePerson lalu PrintPerson. Tangani panic dengan menggunakan recover untuk mencetak pesan error.

Pada kode yang saya buat merupakan program sederhana dalam bahasa Go yang terdiri dari beberapa fungsi utama. 
- Terdapat struct Person yang memiliki dua atribut: Name (string) dan Age (int).
- Fungsi CreatePerson bertugas membuat instance dari struct Person berdasarkan parameter nama dan umur yang diterima. Jika umur yang diberikan kurang dari 0, fungsi ini akan memicu panic dengan pesan "Invalid age".
- Fungsi berikutnya, PrintPerson, digunakan untuk mencetak informasi Person dan dilengkapi dengan mekanisme defer yang menangani potensi panic menggunakan recover. Jika ada panic yang terjadi, pesan yang dihasilkan akan ditampilkan, namun tetap mencetak informasi person jika tidak terjadi panic.
- Selanjutnya, fungsi HandlePerson bertanggung jawab untuk memanggil CreatePerson dan PrintPerson. Fungsi ini juga menggunakan defer dan recover untuk menangani panic dari CreatePerson, sehingga jika panic terjadi, pesan error akan ditampilkan. Pada fungsi main, contoh kasus dijalankan dengan dua pemanggilan HandlePerson. Pemanggilan pertama menggunakan nama "Azka" dan umur -21, yang akan memicu panic karena umur negatif. Pemanggilan kedua menggunakan nama "Caca" dan umur 21, yang akan berhasil mencetak informasi person tanpa error.
