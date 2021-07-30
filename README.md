# SeleksiLabpro-Backend
# Sistem
1. Golang (backend)
2. MYSQL (database)
# Endpoint
**Variant**
----

* GET<br />
  Deskripsi : mengembalikan semua variant dorayaki<br />
  `/variants/`

* GET<br />
  Deskripsi : mengembalikan variant dorayaki berdasarkan id<br />
  `/variants/:id`

* POST<br />
  Deskripsi : menambahkan variant dorayaki baru<br />
  Param : `flavour=[string], description=[string], img_url=[string]`
  `/variants/`

*  PUT<br />
  Deskripsi : memperbarui informasi variant dorayaki berdasarkan id<br />
  Param : `flavour=[string], description=[string], img_url=[string]`
  `/variants/:id`

*  DELETE<br />
  Deskripsi : menghapus  informasi variant dorayaki berdasarkan id<br />
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
1. Pastikan terinstall docker versi > 20
2. Pastikan port `3306` tidak terpakai
3. Jalankan perintah berikut
```
git clone https://github.com/thomas-fm/dorayaki-api.git
cd dorayaki-api
docker-compose up --build
```
4. API akan berjalan pada `localhost:8080/api`

