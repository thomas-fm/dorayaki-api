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

* GET<br />
  Deskripsi : mengembalikan semua toko dorayaki<br />
  `/stores/`

* GET<br />
  Deskripsi : mengembalikan toko dorayaki berdasarkan id<br />
  `/stores/:id`

* POST<br />
  Deskripsi : menambahkan toko dorayaki baru<br />
  Param :`name=[string], street=[string], district=[string], province=[string]`
  `/stores/`

* PUT<br />
  Deskripsi : memperbarui informasi toko dorayaki berdasarkan id<br />
  Param : `name=[string], street=[string], district=[string], province=[string]`
  `/stores/:id`

*  DELETE<br />
  Deskripsi : menghapus toko dorayaki berdasarkan id<br />
  `/stores/:id`

**Stok**
----

* GET<br />
  Deskripsi : mengembalikan semua stok dorayaki pada toko berdasarkan id<br />
  `/stocks/:id`

* GET<br />
  Deskripsi : mengembalikan stok dorayaki dengan id = vId pada toko berdasarkan id<br />
  `/stocks/:id/variant/:vId`

* POST<br />
  Deskripsi : menambahkan stok dorayaki baru pada suatu toko<br />
  Param :`store_id=[integer], variant_id=[integer], total=[integer]`
  `/stocks/:id`

* PUT<br />
  Deskripsi : memperbarui stok dorayaki berdasarkan id toko dan id dorayaki<br />
  Param : `store_id=[integer], variant_id=[integer], total=[integer]`
  `/stocks/:id/variant/:id`

*  DELETE<br />
  Deskripsi : menghapus stok dorayaki berdasarkan id toko dan id dorayaki<br />
  `/stocks/:id/variant/:id`

# Cara menjalankan aplikasi
1. Pastikan terinstall golang versi > 1.16 dan docker versi > 20
2. Pastikan port `3306` tidak terpakai
3. Jalankan perintah berikut<br />
  git clone https://github.com/thomas-fm/dorayaki-api.git<br />
	cd dorayaki-api<br />
	docker-compose up --build<br />
4. API akan berjalan pada `localhost:8080/api`

