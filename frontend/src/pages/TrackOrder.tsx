import { useNavigate } from 'react-router-dom';
import '../styles/pages/TrackOrder.scss'

const TrackOrder = () => {
  const navigate = useNavigate();
  return (
    <section className='track-order'>
      <p>Спасибо за ваш заказ!</p>
      <p>Номер заказа: <strong>9991</strong></p>
      <p>Статус заказа: <strong>Обслуживание</strong></p>
      <button onClick={() => navigate('/')}>Вернуться в магазин</button>
    </section>
  )
}

export default TrackOrder
