import { useParams } from "react-router-dom";
import "../styles/pages/ProductDetails.scss";
import { useCart } from "../hooks/useCart.ts";
import { Product } from "../types";
import { useEffect, useState } from "react";
import { getProduct } from "../api";

const ProductDetails = () => {
  const { id } = useParams<{ id: string }>();
  const { add } = useCart();
  const [product, setProduct] = useState<Product | null>(null);

  useEffect(() => {
    if (id) {
      getProduct(Number(id)).then(setProduct).catch(console.error);
    }
  }, [id])

  if (!product) return ( <div className="product-page">Товар не найден</div>) ;

  const handleAddToCart = () => {
    add({
      id: product.id,
      name: product.name,
      price: product.price,
      image_url: product.image_url
    }, 1);
  };

  return (
    <main className="product-page">
      <div className="product-page__content">
        <img
          className="product-page__image"
          src={product.image_url}
          alt={product.name}
        />
        <div className="product-page__info">
          <h1 className="product-page__title">{product.name}</h1>
          <p className="product-page__price">{product.price}&#8381;</p>
          <div className="product-page__specs">
            <h2 className="product-page__specs-title">Характеристики:</h2>
            <ul className="product-page__specs-list">
              <li className="product-page__specs-item"><p className="product-page__specs-text"><strong>Мощность:</strong> {product.power}</p></li>
              <li className="product-page__specs-item"><p className="product-page__specs-text"><strong>Темертура:</strong> {product.temperature}</p></li>
              <li className="product-page__specs-item"><p className="product-page__specs-text"><strong>Тип:</strong> {product.type}</p></li>
              <li className="product-page__specs-item"><p className="product-page__specs-text"><strong>Цоколь:</strong> {product.lamp_cap}</p></li>
            </ul>
          </div>
          <div className="product-page__description">
            <h2>Описание:</h2>
            <p className="product-page__text">{product.description.String}</p>
          </div>
          <div className="product-page__btn-wrapper">
            <button className="product-page__btn" onClick={handleAddToCart}>Добавить товар в корзину</button>
          </div>
        </div>
      </div>
    </main>
  );
};

export default ProductDetails;
