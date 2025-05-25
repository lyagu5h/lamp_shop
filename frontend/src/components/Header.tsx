import '../styles/components/Header.scss'
import { useCart } from'../hooks/useCart'
import { useNavigate } from 'react-router-dom';

const Header = () => {
  // const { cart } = useCart();

  // useEffect(
  //   () => {
  //     const cart = localStorage.getItem('cart');
  //     if (cart) {
  //       setCart(JSON.parse(cart));
  //     }
  //   },
  //   []
  // )
  const navigate = useNavigate();
  return (
    <header className='header'>
      <h1 className="header__logo">lamp</h1>
      <aside className='header__menu'>
        <button onClick={() => navigate('/')}>Каталог</button>
        <button onClick={() => navigate('/cart')}>Корзина {useCart().cart.length}</button>
      </aside>
    </header>
  )
}

export default Header
