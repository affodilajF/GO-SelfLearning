# ARRAY

- Fixed nilainya, gabisa di append.

- How to make?
  -  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/b17a1cbd-c2ab-4a84-b293-ac1642e5e7e2)
  -  Or you can change [...] to exact number of array.
  -  ![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/157a343a-7538-4d94-a322-1b4396a80791)

# SLICE

- Kinda dynamic version of array (ukurannya bs berubah)
- Slice and Array ALWAYS CONNECTED? what?

- Slice is just to refer data's memorry stored in array.
- Means if theres an slice so theres an array too. 

- The array itself can be explicity written, and the slice referring by [start, stop, step] keyword.
- The aray can be inexplicity written by programmer, it done by the system, the state is achieved by directly writing the slice variable



# INTERFACE

Golang IS NOT OOP lang based. 

In OOP theres a peak parent in which considered as all data implementation. Example : java.lang.Object. 

To achieve the same state in golang, we can utilize blank interface. 

Blank Interface can hold values of any type. why? 

Blank Interface do not have declaration method. Not one. So that, in automatically way, every data type is considered to have implemented the blank interface.

Nama alias dari interface kosong yaitu any. 

interface{} == type any. 

# NIL (data kosong)

Di bbrp bahasa pemograman ada null, di golang ada nil. 

(Tp remember bahwa kalo in another lang, if we make an object without value given, will be considered as null, unlike in go, it will be given default value).

NIL only aplicable for : interface, function, map, slice, pointer, channel. 


# TYPE ASSERTIONS 
The ability to change data type into another data type. 

This feature is often used when along with blank interface. 

The return type is intarface{} but we want its return to be casted as string so here we go. 

But hati-hati, result.(int) is wrong logical syntax though the complier will not say its wrong. 

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/378b01d6-9c6d-4e68-b533-19ecbd587055)

Suggestions : use switch to handle panic. 

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/6bba2dda-06c4-4256-9f45-f3071240b674)

# POINTER (pass by reference)
The ability to refer the memory of exact data without duplicating its value. 

AAAAAAAA I LIKE THE CONCEPT OF POINTER SO MUCH MUCH MUCH MUCHHHHHHH aplg kalo dipake buat parameter di fungsi dan method, you d gonna write so much line if its java!. 

Data type => *data 

Pointer Value => &ArrayReference

# ASTERISTIK OPERATOR (operator bintang) 

### kalo WITHOUT ARTERISTIK OPERATOR : 
Remember, kalo sebuah variabel pointer di assign sama data reference BARU, maka pointer data referensi yang sebelumnya di pointer akan lepas. 

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/5481ef26-cb46-428a-8fca-6b757e8d1f9b)

Hasilnya adalah : 

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/b136e82e-469e-49a7-b3ce-6b1ff1d870d9)

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/1dd1b17a-f097-4827-8bda-21c8aea04647)

## kalo WITH ARTERISTIK OPERATOR 

Karena yg diubah adalah VALUENYA AAAAAAAA okkk.

![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/1915ef01-f711-47ec-aeda-0f4c9a6c8dc8)


![image](https://github.com/affodilajF/GO-SelfLearning/assets/130672181/7a593741-a018-4cd2-88d3-293642b64ca1)


# OPERATOR NEW 
Untuk membuat pointer baru.

Hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal. 

# PACKAGE INITIALIZATION 

fun init() otomatis berjalan ketika sebuah package diimport oleh tempat lain (import "xxx"). 

what for? koneksi database misalnya. 


# PANIC

# DEFER


