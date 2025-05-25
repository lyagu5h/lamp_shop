import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../hooks/useCart";
import "../styles/pages/Cart.scss";
import CartModal from "../components/CartModal";

const CartPage = () => {
  const navigate = useNavigate();
  const { cart, error, calculateTotal } = useCart();

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
                <span>{calculateTotal().toFixed(2)} ₽</span>
              </div>
            </div>
          </div>

          <form className="order-form">

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
                  />
                </div>
              </div>

            <div className="order-form__group">
              <label className="order-form__section-title order-form__label--hidden" htmlFor="comment">
                Комментарий к заказу
              </label>
              <textarea
                className="order-form__comment"
                id="comment"
                name="comment"
                rows={4}
                placeholder="Ваши пожелания..."
              ></textarea>
            </div>

            <button className="order-form__submit" type="button" onClick={() => navigate('/orders/242342/status')}>
              Оформить заказ
            </button>
          </form>
        </>
      )}
    </div>
  );
};

export default CartPage;
