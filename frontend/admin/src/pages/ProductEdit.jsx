import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  fetchProductById,
  createProduct,
  updateProduct,
  uploadProductImage,
} from "../api/products";

export default function ProductEdit() {
  const { id } = useParams();       
  const isNew = !id;
  const navigate = useNavigate();

  const [form, setForm] = useState({
    name: "",
    price: 0,
    power: 0,
    description: {
        String: "",
        Valid: false
    },
    temperature: "",
    type: "",
    lamp_cap: "",
  });

  const [currentImageUrl, setCurrentImageUrl] = useState(null);

  const [file, setFile] = useState(null);

  useEffect(() => {
    if (!isNew) {
      fetchProductById(id)
        .then((data) => {
          setForm({
            name: data.name || "",
            price: Number(data.price) || 0,
            power: Number(data.power) || 0,
            description: {
                String: data.description.String || "",
                Valid: data.description.Valid || false
            },
            temperature: data.temperature || "",
            type: data.type || "",
            lamp_cap: data.lamp_cap || "",
          });
          setCurrentImageUrl(data.image_url || null);
        })
        .catch((err) => {
          console.error(err);
          navigate("/products");
        });
    }
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name === "price" || name === "power") {
      setForm((prev) => ({ ...prev, [name]: Number(value) }));
    } else if (name === "description") {
      setForm((prev) => ({
        ...prev,
        description: {
          String: value,
          Valid: Boolean(value),
        },
      }));
    } else {
      setForm((prev) => ({ ...prev, [name]: value }));
    }
  };

  const handleFileChange = (e) => {
    if (e.target.files && e.target.files.length > 0) {
      console.log("Selected file:", e.target.files[0]);
      setFile(e.target.files[0]);
    } else {
      setFile(null);
    }
  };

  const uploadImage = async (productId) => {
    if (!file) return null;
    const formData = new FormData();
    formData.append("image", file);

    for (let pair of formData.entries()) {
      console.log(pair[0], pair[1]);
    }
    try {
      const response = await uploadProductImage(productId, formData);
      return response.image_url || null;
    } catch (err) {
      console.error("Ошибка при загрузке изображения:", err);
      return null;
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let productId = id;

      if (isNew) {
        const created = await createProduct(form);
        productId = created.id;
      } else {
        await updateProduct(id, form);
      }

      if (file) {
        const newUrl = await uploadImage(productId);
        if (newUrl) {
          console.log("Картинка загружена:", newUrl);
        }
      }

      navigate("/products");
    } catch (err) {
      console.error(err);
      alert("Ошибка при сохранении товара");
    }
  };

  return (
    <div className="max-w-xl mx-auto bg-white p-6 rounded shadow mt-8">
      <h2 className="text-2xl font-medium mb-4">
        {isNew ? "Создать новый товар" : `Редактировать товар #${id}`}
      </h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Название */}
        <div>
          <label className="block mb-1 font-semibold">Название</label>
          <input
            name="name"
            type="text"
            className="w-full border px-3 py-2 rounded"
            placeholder="Лампа A60 LED"
            value={form.name}
            onChange={handleChange}
            required
          />
        </div>

        {/* Цена */}
        <div>
          <label className="block mb-1 font-semibold">Цена (₽)</label>
          <input
            name="price"
            type="number"
            step="0.01"
            className="w-full border px-3 py-2 rounded"
            placeholder="Введите цену"
            value={form.price}
            onChange={handleChange}
            required
          />
        </div>

        {/* Мощность */}
        <div>
          <label className="block mb-1 font-semibold">Мощность (Вт)</label>
          <input
            name="power"
            type="number"
            className="w-full border px-3 py-2 rounded"
            placeholder="Например: 7"
            value={form.power}
            onChange={handleChange}
          />
        </div>

        {/* Описание */}
        <div>
          <label className="block mb-1 font-semibold">Описание</label>
          <textarea
            name="description"
            className="w-full border px-3 py-2 rounded"
            placeholder="Краткое описание товара"
            rows={3}
            value={form.description.String}
            onChange={handleChange}
          />
        </div>

        {/* Температура */}
        <div>
          <label className="block mb-1 font-semibold">Цветовая температура (K)</label>
          <input
            name="temperature"
            type="text"
            className="w-full border px-3 py-2 rounded"
            placeholder="Например: 2700"
            value={form.temperature}
            onChange={handleChange}
          />
        </div>

        {/* Тип лампы */}
        <div>
          <label className="block mb-1 font-semibold">Тип лампы</label>
          <input
            name="type"
            type="text"
            className="w-full border px-3 py-2 rounded"
            placeholder="Например: LED"
            value={form.type}
            onChange={handleChange}
          />
        </div>

        {/* Цоколь */}
        <div>
          <label className="block mb-1 font-semibold">Цоколь</label>
          <input
            name="lamp_cap"
            type="text"
            className="w-full border px-3 py-2 rounded"
            placeholder="Например: E27"
            value={form.lamp_cap}
            onChange={handleChange}
          />
        </div>

        {/* Текущее изображение (если редактируем) */}
        {currentImageUrl && (
          <div>
            <label className="block mb-1 font-semibold">Текущее изображение</label>
            <img
              src={currentImageUrl}
              alt="Текущее изображение"
              className="w-32 h-32 object-cover rounded border"
            />
          </div>
        )}

        {/* Выбор нового файла */}
        <div>
          <label className="block mb-1 font-semibold">Изображение</label>
          <input
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="block"
          />
          <p className="text-sm text-gray-500 mt-1">
            {isNew
              ? "Выберите файл для загрузки изображения"
              : "Выберите файл, чтобы заменить текущее изображение"}
          </p>
        </div>

        {/* Дата создания/обновления (только для просмотра) */}
        {!isNew && (
          <div className="text-xs text-gray-500">
            <p>
              <strong>Создано:</strong>{" "}
              {new Date(form.created_at).toLocaleString()}
            </p>
            <p>
              <strong>Обновлено:</strong>{" "}
              {new Date(form.updated_at).toLocaleString()}
            </p>
          </div>
        )}

        {/* Кнопки */}
        <div className="flex justify-end space-x-2 pt-4">
          <button
            type="button"
            onClick={() => navigate("/products")}
            className="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
          >
            Отмена
          </button>
          <button
            type="submit"
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
          >
            Сохранить
          </button>
        </div>
      </form>
    </div>
  );
}
