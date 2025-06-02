import '../styles/components/Header.scss'
import { useCart } from'../hooks/useCart.ts'
import { useNavigate } from 'react-router-dom';

const Header = () => {
  const { cart } = useCart();
  const navigate = useNavigate();

  return (
    <header className='header'>
      <h1 className="header__logo">lamp</h1>
      <aside className='header__menu'>
        <button onClick={() => navigate('/')}>Каталог</button>
        <button onClick={() => navigate('/cart')}>Корзина {cart.length}</button>
      </aside>
    </header>
  )
}

export default Header
