# Aplikasi To-Do List CLI
### Mini Project App To-Do List CLI oleh Fathoni

**Video Demo:** [`  Video-Demo-To-Do-List-CLI   `](https://youtu.be/nZkQ6N-Ba7o)

## Tentang Aplikasi
Aplikasi To-Do List CLI merupakan aplikasi pembuatan jadwal dengan fitur menambah, menampilkan seluruh task, mencari, memperbarui, dan menghapus task.

**Add Task**
````bash
Go run . add -t="task name" -p="medium"
````
````bash
-t      atau    --task          -> [masukkan nama task]
-p      atau    --priority      -> [masukkan priority task: low/medium/high]
````

**List Task**
````bash
Go run . list
````

**Search Task**
````bash
Go run . search -t="task name"
````
````bash
-t      atau    --task          -> [masukkan nama task]
````

**Update Task**
````bash
Go run . update -i=1 -t="task name" -s="progress" -p="medium"
````
````bash
-i      atau    --id            -> [masukkan id task]
-t      atau    --task          -> [masukkan nama task]
-s      atau    --status        -> [masukkan status task: new/pending/progress/completed]
-p      atau    --priority      -> [masukkan priority task: low/medium/high]
````

**Delete Task**
````bash
Go run . delete -i=1
````
````bash
-i      atau    --id            -> [masukkan id task]
````

Terima kasih telah menggunakan Aplikasi To-Do List CLI
`Happy Coding ❤️🩵`