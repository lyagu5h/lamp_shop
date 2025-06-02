import { productsApi } from "./axios";


export async function fetchProducts() {
  const res = await productsApi.get("/");
  return res.data;
}


export async function fetchProductById(id) {
  const res = await productsApi.get(`/${id}`);
  return res.data;
}


export async function createProduct(data) {
  const res = await productsApi.post("/", data);
  return res.data;
}


export async function updateProduct(id, data) {
  console.log(typeof data.power)
  const res = await productsApi.put(`/${id}`, data);
  return res.data;
}


export async function deleteProduct(id) {
  const res = await productsApi.delete(`/${id}`);
  return res.data;
}


export async function uploadProductImage(id, formData) {
  const res = await productsApi.post(`/${id}/image`, formData);
  return res.data;
}