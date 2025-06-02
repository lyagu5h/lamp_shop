import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  fetchProducts,
  deleteProduct
} from "../api/products";

// Модель Product (для подсказки, не обязательно импортируется):
// type Product = {
//   id: number;
//   name: string;
//   price: number;
//   power: number;
//   description: string | null;
//   temperature: string;
//   type: string;
//   lamp_cap: string;
//   image_url: string;
//   created_at: string; // ISO
//   updated_at: string; // ISO
// };

export default function ProductsList() {
  const [products, setProducts] = useState([]);
  const navigate = useNavigate();

  const loadProducts = async () => {
    try {
      const data = await fetchProducts(); // ожидаем array of Product
      setProducts(data);
    } catch (err) {
      console.error(err);
      // Если, например, 401, редиректим на /login
      navigate("/login");
    }
  };

  useEffect(() => {
    loadProducts();
  }, []);

  const handleDelete = async (id) => {
    if (!window.confirm("Удалить этот товар?")) return;
    try {
      await deleteProduct(id);
      loadProducts();
    } catch (err) {
      console.error(err);
      alert("Ошибка при удалении товара");
    }
  };

  return (
    <div className="px-4 py-6">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-semibold">Управление товарами</h2>
        <button
          onClick={() => navigate("/products/new")}
          className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
        >
          Добавить товар
        </button>
      </div>

      <div className="overflow-x-auto bg-white rounded shadow">
        <table className="min-w-full text-left text-sm">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-3 py-2">ID</th>
              <th className="px-3 py-2">Название</th>
              <th className="px-3 py-2">Цена (₽)</th>
              <th className="px-3 py-2">Мощность</th>
              <th className="px-3 py-2">Описание</th>
              <th className="px-3 py-2">Темп. (K)</th>
              <th className="px-3 py-2">Тип лампы</th>
              <th className="px-3 py-2">Цоколь</th>
              <th className="px-3 py-2">Картинка</th>
              <th className="px-3 py-2">Создан</th>
              <th className="px-3 py-2">Обновлён</th>
              <th className="px-3 py-2">Действия</th>
            </tr>
          </thead>
          <tbody>
            {products.map((p) => (
              <tr key={p.id} className="border-b hover:bg-gray-50">
                <td className="px-3 py-2">{p.id}</td>
                <td className="px-3 py-2">{p.name}</td>
                <td className="px-3 py-2">{p.price.toFixed(2)}</td>
                <td className="px-3 py-2">{p.power}</td>
                <td className="px-3 py-2">
                  <div className="max-w-xs truncate">{p.description.String || "-"}</div>
                </td>
                <td className="px-3 py-2">{p.temperature}</td>
                <td className="px-3 py-2">{p.type}</td>
                <td className="px-3 py-2">{p.lamp_cap}</td>
                <td className="px-3 py-2">
                  {p.image_url ? (
                    <img
                      src={p.image_url}
                      alt={p.name}
                      className="w-12 h-12 object-cover rounded"
                    />
                  ) : (
                    "-"
                  )}
                </td>
                <td className="px-3 py-2 text-xs text-gray-500">
                  {new Date(p.created_at).toLocaleString()}
                </td>
                <td className="px-3 py-2 text-xs text-gray-500">
                  {new Date(p.updated_at).toLocaleString()}
                </td>
                <td className="px-3 py-2 flex space-x-2">
                  <button
                    onClick={() => navigate(`/products/${p.id}/edit`)}
                    className="bg-blue-500 hover:bg-blue-600 text-white px-2 py-1 rounded text-xs"
                  >
                    ✏️ Ред
                  </button>
                  <button
                    onClick={() => handleDelete(p.id)}
                    className="bg-red-500 hover:bg-red-600 text-white px-2 py-1 rounded text-xs"
                  >
                    🗑️ Удалить
                  </button>
                </td>
              </tr>
            ))}

            {products.length === 0 && (
              <tr>
                <td colSpan="12" className="px-3 py-4 text-center text-gray-500">
                  Нет товаров
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
