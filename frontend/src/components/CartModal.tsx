import { useCart, CartItem } from "../hooks/useCart"
import '../styles/components/CartModal.scss'

const CartModal = (item: CartItem) => {
  const { removeFromCart, updateQuantity } = useCart()
  return (
    <div key={item.id} className="cart-item">
      <div className="cart-item__container">
        <div className="cart-item__img-container">
          <img 
            src={item.image_url} 
            alt={item.name}
            className="cart-item__img"
          />
        </div>
        <button 
            onClick={() => removeFromCart(item.id)}
            className="cart-item__remove-btn"
          >
            Удалить
          </button>
          
      </div>

      
      <div className="cart-item__info">
        <h3 className="cart-item__name">{item.name}</h3>
        <p className="cart-item__price">{item.price} ₽</p>
        <div className="cart-item__quantity-container">
            <button 
              onClick={() => updateQuantity(item.id, item.quantity - 1)}
              disabled={item.quantity <= 1}
              className="cart-item__quantity-btn"
            >
              -
            </button>
            <span className="cart-item__quantity-value">{item.quantity}x</span>
            <button 
              onClick={() => updateQuantity(item.id, item.quantity + 1)}
              className="cart-item__quantity-btn"
            >
              +
            </button>          
        </div>
        

      </div>
    </div>
  )
}

export default CartModal
