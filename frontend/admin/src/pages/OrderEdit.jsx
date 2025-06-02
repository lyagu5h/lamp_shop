import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { fetchOrderById, updateOrderStatus } from "../api/orders";

export default function OrderEdit() {
  const { id } = useParams(); // ID заказа
  const navigate = useNavigate();

  const [order, setOrder] = useState(null);
  const [newStatus, setNewStatus] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchOrderById(id)
      .then((data) => {
        setOrder(data);
        setNewStatus(data.status);
      })
      .catch((err) => {
        console.error(err);
        navigate("/orders");
      })
      .finally(() => {
        setLoading(false);
      });
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (newStatus === order.status) {
      navigate("/orders");
      return;
    }
    try {
      await updateOrderStatus(id, {status: newStatus});
      navigate("/orders");
    } catch (err) {
      console.error(err);
      alert("Ошибка при обновлении статуса заказа");
    }
  };

  if (loading) {
    return <p className="p-4 text-gray-600">Загрузка данных заказа...</p>;
  }
  if (!order) {
    return null;
  }

  return (
    <div className="max-w-lg mx-auto bg-white p-6 rounded shadow mt-8">
      <h2 className="text-2xl font-semibold mb-4">Заказ #{order.id}</h2>

      <div className="space-y-2 mb-6">
        <p>
          <strong>Имя клиента:</strong> {order.customer_name}
        </p>
        <p>
          <strong>Email:</strong> {order.customer_email}
        </p>
        <p>
          <strong>Телефон:</strong> {order.customer_phone}
        </p>
        <p>
          <strong>Адрес:</strong> {order.address}
        </p>
        <p>
          <strong>Сумма (₽):</strong> {order.total_amount.toFixed(2)}
        </p>
        <p>
          <strong>Текущий статус:</strong> {order.status}
        </p>
        <p className="text-xs text-gray-500">
          Создан: {new Date(order.created_at).toLocaleString()}
        </p>
        <p className="text-xs text-gray-500">
          Обновлён: {new Date(order.updated_at).toLocaleString()}
        </p>
      </div>

      <div className="mb-6">
        <p className="font-semibold mb-2">Товары в заказе:</p>
        <ul className="list-disc list-inside space-y-1 text-sm">
          {order.items.map((item) => (
            <li key={item.id}>
              Продукт #{item.product_id} — количество: {item.quantity}, цена:{" "}
              {item.price.toFixed(2)} ₽
            </li>
          ))}
        </ul>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block mb-1 font-semibold">Изменить статус</label>
          <select
            value={newStatus}
            onChange={(e) => setNewStatus(e.target.value)}
            className="w-full border px-3 py-2 rounded"
          >
            <option value="new">new</option>
            <option value="processed">processed</option>
            <option value="shipped">shipped</option>
            <option value="delivered">delivered</option>
            <option value="cancelled">cancelled</option>
          </select>
        </div>

        <div className="flex justify-end space-x-2 pt-4">
          <button
            type="button"
            onClick={() => navigate("/orders")}
            className="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
          >
            Отмена
          </button>
          <button
            type="submit"
            className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded"
          >
            Сохранить
          </button>
        </div>
      </form>
    </div>
  );
}
