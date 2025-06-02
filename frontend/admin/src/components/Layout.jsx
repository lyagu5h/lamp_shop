import React from "react";
import { Link, useNavigate, Outlet } from "react-router-dom";

export default function Layout() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-black text-white p-4 flex justify-between items-center">
        <div className="flex space-x-4">
          <Link to="/products" className="hover:underline">
            Товары
          </Link>
          <Link to="/orders" className="hover:underline">
            Заказы
          </Link>
        </div>
        <button onClick={handleLogout} className="text-red-400 hover:underline">
          Выход
        </button>
      </nav>

      <main className="p-4">
        <Outlet />
      </main>
    </div>
  );
}