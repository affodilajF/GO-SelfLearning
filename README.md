
# GOROUTINE 
- Goroutine berjalan didalam thread, jumlah thread diatur sama GOMAXPROCS (biasanya sejumlah core cpu)
- Goroutine is often to be called "light thread" just bc the size is only 1kb (thread having 1mb). Itu gpp.
- Whereas goroutine is runinng inside thread. 
- Goroutine diatur sama Go-Schduler

- P is core, M is thread , oooo itu adalah goroutine.
- ![Screenshot 2024-04-06 112057](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/5e710e5b-c620-4bd4-b76e-444b17f511f8)


# CHANNEL 
- Goroutine cannot return data bc of complexity reason (it will make the coding activity harder)
- In order to get data from other goroutine, we can utilize channel
- Channel is such a communication bridge between 2 goroutines
- Channel is storing data
- The data can be send or receive through goroutine

![Screenshot 2024-04-06 114153](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/234d6394-e8cf-4170-ae6c-d0f20874eb30)

#### Channel characteristics 
- Only acc one type of data. And its need to be defined before.
- Only accept one data. We need to wait until the first data is picked by the goroutine before passing the second data.
- Bc of memory leak haunting, channel must be close.
- Channel bisa diambil dari LEBIH SATU GOROUTINE. what does it mean?
  
## Make Channel 
- ![Screenshot 2024-04-06 115657](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/4477119f-267e-4907-b5f7-ca308141fc14)

## Sending and Receiving data from channel 
- Sending : chan <- data
- Receiving : data <- chan
- Closing : close(chanVarName)
- ![Screenshot 2024-04-06 123335](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/c2dec87c-a2ac-4aa9-b6c0-038c914db3c3)
- ![Screenshot 2024-04-06 124007](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/c8aabd7c-8665-4edb-ac46-b1330d0fe219)

## Channel as parameter
- By default, channel is pass by referende. SO we dont need pointer.
- ![Screenshot 2024-04-06 130201](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/708986bd-b5f3-4407-a862-1bd2995fec5a)

## Channel In and Out 
- In channel as parameter, the function can either sending or receiving the data form its channel.
- But we can explicitly inform the function that the channel is intended only for either receiving or sending data
- ![Screenshot 2024-04-06 130649](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/8173141e-4bf9-4de9-97ac-5ea849990e52)

## Buferred Channel (penyimpanan antrian di channel)
- Sending is faster than receiving. Wdym?
- For instance, we have 2 activity to send-receive a data from 1 channel.

  If its ordinary channel, and if sending data in 2nd processes is faster than receiving data in 1st processes, the 2nd sending processes will be ikutan lambat juga.
  Tackle? use buffered channel!
- Buffer is to accomondate quequed data in channel.

#### Buffer Capacity 
- Bebas berapa kapasitasnya
- Misal di set kapasitasnya 4 ya brti kalo ada goro ke 5 harus nunggu kapasitas buffernya selo 1

- Misal ni channel di set kapasitasnya goroutine maka :
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/9d5f9e91-bd50-4e86-a57f-0d830274687f)

- To make?
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/23522d37-6dc2-4d90-a60d-d57482921683)

- GINI

  Pake buffer channel bisa mencegah deadlock.
  Deadlock terjadi ketika dua atau lebih entitas (biasanya proses atau goroutine) dalam sebuah sistem terjebak dalam siklus di mana masing-masing menunggu sumber daya yang dikendalikan oleh yang lain.
  Misalnya pake channel tapi cuman ada goroutine pengirim doang, goro receivernya gaada, ntar bakalan blocking lama trs nyebabin deadlock.

  Pake buffer capacity!. Misal masukkin data 2x kedalam channel, data pertama belum diambil sama goro (bisa karena receive nya lama), data ke-2 bakalan ada di queque buffer.

  Tapi, kapasitasnya 4, dan buffer channelnya full, trs kamu nambahin 1 data lagi, bakalan tetap deadlock.


## Range Channel 
- Theres an case, sebuah channel dikirim data terus2 an oleh goro sender. And we cant make sure when it will be ended.
- Salah satu solusi : use perulangan range ketika menerima data dari channel.
- When close() is invoked, automatically the loop is stopped.
- This is simpler than manually checking the channel.
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/35f58cac-ab2f-4c68-92b5-bdd51769c293)

## Select Channel
- Theres a case, ketika kita membuat beberapa channel dan menjalankan beberapa goroutine. Lalu kita ingin mendaptkan semua data dari channel tsb.
- Untuk melakukan ini, kita bisa menggunakan select channel di Go-Lang.

- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel. Jika data datang bersamaan di beberapa channel, maka akan dipilih secara random.
- 1 select cuma bisa pilih 1 channel, by default.
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/24ed383d-e704-4f64-a977-017fc1049a11)
- Biar bisa ngambil> 1 data, pakai for. Tapi jangan lupa loopnya di stop, ntar deadlock.
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/6c30cade-2953-4e96-adff-0e83168f7f04)

#### Default Select 
- If theres no data inside channel but you have made select channel statement, it will result in a deadlock.
- DEFAULT statement will be executed if all the channel have no data (in possesion of waiting the data). 
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/9540c8a8-402a-44bf-a7a5-44fc7066b9e6)

  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/637556b8-eb6a-4a63-9e12-70f2009d0ebd)

# Race Condition 
- Penyebab => sharing variabel yang diakses dan diubah oleh banyak goroutine.

  Karena ada paralel, nilai x bisa diakses 2x secara bersamaan dan paralel. Balapan mereka. 
  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/844cb9ad-72be-4e69-8826-a0325689d5eb)

### CARA MENGATASI RACE CONDITION : 
### sync Mutex 
- Adalah sebuah struct
- To locking and unlocking
- If several gotoutines try to lock a Mutex simultaneously, only ONE goroutine will be allowed. Once the goroutine that acquired the lock has finished its operation and unlock the Mutex, the next goroutine is allowed to lock it.
- ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/d5e90d31-4bca-491b-8991-28cb46a94b18)

### sync RWMutex (Read Write Mutex)
- Have 2 lock => Read lock and Write lock
- ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/b38a7a22-db08-41ad-b9cb-176bab4ee58c)

# Deadlock 
- Hati-hati pas bikin aplikasi yang paralel atau concurrent!
- Deadlock occurs when goroutine processes are waiting for each other's locks, preventing any goroutine from running.
  
### sync WaitGroup (nah ini dipake in real project)
- Go's feature, allows to wait concurent processes to be done.
- Goroutine will be executed and waited for until all of them have finished. 
- Biasanya km pake thread sleep kan, tp itu ngira2 brp detik, jlek, mending pake ini yaaa.
- Add(int) method => to mark theres an goroutine processes
- Done() method => after goroutine processes have been finished/done.
- Wait() method => to wait all the processes done.
- ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/d26b720f-5211-469a-84c5-efb914a67f42)
- 

### sync.Once 
- To make sure exact function is only once executed
- Berapapun goro yg mau coba akses, cuma yang awal doang yang boleh
- ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/d6a67221-1465-4e47-9806-5a6f7416406a)

## sync Pool 
- implementation of design pattern Object Pool Pattern (emg d khususkan untuk concurent and paralel)
- DesPat ini untuk menyimpan data, selanjutnya untuk menggambil datanya bisa dari Pool, setelah slesai menggunakan datanya, balikin lagi ke Poolnya
- This implementation is far from being threatened by race conditions
- Biasanya dipake buat koneksi db (biar ga berkali2 konek ke db, mahal)
- put => masukkin data

  get => ambil data
  
   ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/2ad56d41-fbcd-424f-83d9-400401269d57)

  
  
## sync Map 
## sync Cond
## atomic
## time Timer
## time Ticker
## GOMAXPROCS

  





