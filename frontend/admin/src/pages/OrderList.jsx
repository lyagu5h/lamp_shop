import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { fetchOrders } from "../api/orders";
import { fetchProducts } from "../api/products";


export default function OrdersList() {
  const [orders, setOrders] = useState([]);
  const [products, setProducts] = useState({});
  const navigate = useNavigate();

  const loadOrders = async () => {
    try {
      const data = await fetchOrders();
      setOrders(data);
    } catch (err) {
      console.error(err);
      navigate("/login");
    }
  };

  const loadProducts = async() => {
    try {
      fetchProducts().then(prods => {
        const m = Object.fromEntries(prods.map(p => [p.id, p]));
        setProducts(m);
      });
    } catch (err) {
      console.error(err);
      navigate("/login");
    }
  }

  useEffect(() => {
    loadOrders();
    loadProducts();
  }, []);

  return (
    <div className="px-4 py-6">
      <h2 className="text-2xl font-semibold mb-4">Список заказов</h2>

      <div className="overflow-x-auto bg-white rounded shadow">
        <table className="min-w-full text-left text-sm">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-3 py-2">ID</th>
              <th className="px-3 py-2">Имя клиента</th>
              <th className="px-3 py-2">Email</th>
              <th className="px-3 py-2">Телефон</th>
              <th className="px-3 py-2">Адрес</th>
              <th className="px-3 py-2">Сумма (₽)</th>
              <th className="px-3 py-2">Статус</th>
              <th className="px-3 py-2">Дата создания</th>
              <th className="px-3 py-2">Действия</th>
            </tr>
          </thead>
          <tbody>
            {orders.map((o) => (
              <React.Fragment key={o.id}>
                <tr className="border-b hover:bg-gray-50">
                  <td className="px-3 py-2">{o.id}</td>
                  <td className="px-3 py-2">{o.customer_name}</td>
                  <td className="px-3 py-2">{o.customer_email}</td>
                  <td className="px-3 py-2">{o.customer_phone}</td>
                  <td className="px-3 py-2">
                    <div className="max-w-sm truncate">{o.address}</div>
                  </td>
                  <td className="px-3 py-2">{o.total_amount.toFixed(2)}</td>
                  <td className="px-3 py-2">
                    <span
                      className={`px-2 py-1 rounded text-xs font-medium ${
                        o.status === "new"
                          ? "bg-yellow-100 text-yellow-800"
                          : o.status === "processed"
                          ? "bg-green-100 text-green-800"
                          : o.status === "shipped"
                          ? "bg-blue-100 text-blue-800"
                          : "bg-gray-100 text-gray-800"
                      }`}
                    >
                      {o.status}
                    </span>
                  </td>
                  <td className="px-3 py-2 text-xs text-gray-500">
                    {new Date(o.created_at).toLocaleString()}
                  </td>
                  <td className="px-3 py-2">
                    <button
                      onClick={() => navigate(`/orders/${o.id}`)}
                      className="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded text-xs"
                    >
                      Детали
                    </button>
                  </td>
                </tr>

                <tr className="bg-gray-50">
                  <td colSpan="9" className="px-3 py-2">
                    <div className="p-2 bg-white border rounded">
                      <p className="text-sm font-medium mb-1">Товары в заказе:</p>
                      <ul className="list-disc list-inside text-sm">
                        {/* {o.items.map((item) => (
                          <li key={item.id}>
                            {item.product_id} — количество: {item.quantity}, цена:{" "}
                            {item.price.toFixed(2)} ₽
                          </li>
                        ))} */}
                         {o.items.map((item) => (
                          <li key={item.id}>
                            {products[item.product_id].name} — количество: {item.quantity}, цена:{" "}{item.price.toFixed(2)} ₽
                          </li>
                        ))}
                      </ul>
                    </div>
                  </td>
                </tr>
              </React.Fragment>
            ))}

            {orders.length === 0 && (
              <tr>
                <td colSpan="9" className="px-3 py-4 text-center text-gray-500">
                  Нет заказов
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
