import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";

import PrivateRoute from "./components/PrivateRoute";
import Layout from "./components/Layout";

import Login from "./pages/Login";
import ProductsList from "./pages/ProductList";
import ProductEdit from "./pages/ProductEdit";
import OrdersList from "./pages/OrderList";
import OrderEdit from "./pages/OrderEdit";

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />

        <Route
          path="/"
          element={
            <PrivateRoute>
              <Layout />
            </PrivateRoute>
          }
        >
          <Route index element={<Navigate to="/products" />} />

          <Route path="products" element={<ProductsList />} />
          <Route path="products/:id/edit" element={<ProductEdit />} />
          <Route path="products/new" element={<ProductEdit />} />

          <Route path="orders" element={<OrdersList />} />
          <Route path="orders/:id" element={<OrderEdit />} />
        </Route>

        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </Router>
  );
}