# SeleksiLabpro-Backend

# REST API
**Variant**
----

* GET
  Deskripsi : mengembalikan semua variant dorayaki
  `/variants/`

* GET
  Deskripsi : mengembalikan variant dorayaki berdasarkan id
  `/variants/:id`

* POST
  Deskripsi : menambahkan variant dorayaki baru
  Param : `flavour=[string], description=[string], img_url=[string]`
  `/variants/`

*  PUT
  Deskripsi : memperbarui informasi variant dorayaki berdasarkan id
  Param : `flavour=[string], description=[string], img_url=[string]`
  `/variants/:id`

*  DELETE
  Deskripsi : menghapus  informasi variant dorayaki berdasarkan id
  `/variants/:id`

**Toko**
----

* GET
  Deskripsi : mengembalikan semua toko dorayaki
  `/stores/`

* GET
  Deskripsi : mengembalikan toko dorayaki berdasarkan id
  `/stores/:id`

* POST
  Deskripsi : menambahkan toko dorayaki baru
  Param :`name=[string], street=[string], district=[string], province=[string]`
  `/stores/`

* PUT
  Deskripsi : memperbarui informasi toko dorayaki berdasarkan id
  Param : `name=[string], street=[string], district=[string], province=[string]`
  `/stores/:id`

*  DELETE
  Deskripsi : menghapus toko dorayaki berdasarkan id
  `/stores/:id`

**Stok**
----

* GET
  Deskripsi : mengembalikan semua stok dorayaki pada toko berdasarkan id
  `/stocks/:id`

* GET
  Deskripsi : mengembalikan stok dorayaki dengan id = vId pada toko berdasarkan id
  `/stocks/:id/variant/:vId`

* POST
  Deskripsi : menambahkan stok dorayaki baru pada suatu toko
  Param :`store_id=[integer], variant_id=[integer], total=[integer]`
  `/stocks/:id`

* PUT
  Deskripsi : memperbarui stok dorayaki berdasarkan id toko dan id dorayaki
  Param : `store_id=[integer], variant_id=[integer], total=[integer]`
  `/stocks/:id/variant/:id`

*  DELETE
  Deskripsi : menghapus stok dorayaki berdasarkan id toko dan id dorayaki
  `/stocks/:id/variant/:id`

# Cara menjalankan aplikasi
1. Pastikan terinstall golang versi > 1.16 dan docker versi > 20
2. Pastikan port `3306` tidak terpakai
3. Jalankan perintah berikut
  git clone https://github.com/thomas-fm/dorayaki-api.git
	cd dorayaki-api
	docker-compose up --build
4. API akan berjalan pada `localhost:8080/api`

