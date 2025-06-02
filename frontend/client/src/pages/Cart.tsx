import { useState, FormEvent } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useCart } from '../hooks/useCart.ts';
import { createOrder } from '../api';
import CartModal from '../components/CartModal';
import "../styles/pages/Cart.scss";
import { CartItem } from '../hooks/cartContext';

const CartPage = () => {
    const {
    cart,
    clear,
    total,
  } = useCart();

  const navigate = useNavigate();

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [city, setCity] = useState('');
  const [street, setStreet] = useState('');
  const [house, setHouse] = useState('');

  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    if (!cart.length) return;

    setError(null);

    try {
      const order = await createOrder({
        customer_name : name,
        customer_email: email,
        customer_phone: phone,
        address      : `${city}, ${street}, д.${house}`,
        total_amount : total(),
        items        : cart.map(({ id, quantity, price }: CartItem) => ({
          product_id: id,
          quantity,
          price,
        })),
      });

      clear();                 
      navigate(`/orders/${order.id}/status`);
    } catch (err: unknown) {
      setError((err as Error).message ?? 'Не удалось оформить заказ');
    }
  };

  return (
    <div className="cart-page">
      {error && (
        <div className="cart-page__error">
          <p>Error: {error}</p>
        </div>
      )}

      {cart.length === 0 ? (
        <div className="cart-page__empty">
          <p className="cart-page__empty-text">Ваша корзина пуста</p>
          <Link to="/" className="cart-page__continue-btn">
            Продолжить покупки
          </Link>
        </div>
      ) : (
        <>
          <div className="cart-page__items">
            {cart.map((item) => (
              <CartModal key={item.id} {...item} />
            ))}
            <div className="cart-page__summary">
              <div className="cart-page__total">
                <span>Итого: </span>
                <span>{total().toFixed(2)} ₽</span>
              </div>
            </div>
          </div>

          <form className="order-form" onSubmit={handleSubmit}>

              <div className="order-form__group">
                <label
                  className="order-form__label order-form__label--hidden"
                  htmlFor="email"
                >
                  Имя
                </label>
                <input
                  className="order-form__input"
                  type="name"
                  id="name"
                  name="name"
                  placeholder="Имя"
                  onChange={e => setName(e.target.value)}
                />
              </div>
              
              <div className="order-form__group">
                <label
                  className="order-form__label order-form__label--hidden"
                  htmlFor="email"
                >
                  Email
                </label>
                <input
                  className="order-form__input"
                  type="email"
                  id="email"
                  name="email"
                  placeholder="Email"
                  onChange={e => setEmail(e.target.value)}
                />
              </div>

              <div className="order-form__group">
                <label
                  className="order-form__label order-form__label--hidden"
                  htmlFor="phone"
                >
                  Номер телефона
                </label>
                <input
                  className="order-form__input"
                  type="tel"
                  id="phone"
                  name="phone"
                  placeholder="Номер телефона"
                  onChange={e => setPhone(e.target.value)}
                />
              </div>

              <div className="order-form__group">
                <label
                  className="order-form__label order-form__label--hidden"
                  htmlFor="city"
                >
                  Город
                </label>
                <input
                  className="order-form__input"
                  type="text"
                  id="city"
                  name="city"
                  placeholder="Город"
                  onChange={e => setCity(e.target.value)}
                />
              </div>

              <div className="order-form__row order-form__row--grouped">
                <div className="order-form__group order-form__group--street">
                  <label
                    className="order-form__label order-form__label--hidden"
                    htmlFor="street"
                  >
                    Улица
                  </label>
                  <input
                    className="order-form__input"
                    type="text"
                    id="street"
                    name="street"
                    placeholder="Улица"
                    onChange={e => setStreet(e.target.value)}
                  />
                </div>

                <div className="order-form__group order-form__group--house">
                  <label
                    className="order-form__label order-form__label--hidden"
                    htmlFor="house"
                  >
                    Дом
                  </label>
                  <input
                    className="order-form__input"
                    type="text"
                    id="house"
                    name="house"
                    placeholder="Дом"
                    onChange={e => setHouse(e.target.value)}
                  />
                </div>
              </div>

            <button className="order-form__submit" type="submit">
              Оформить заказ
            </button>
          </form>
        </>
      )}
    </div>
  );
};

export default CartPage;
