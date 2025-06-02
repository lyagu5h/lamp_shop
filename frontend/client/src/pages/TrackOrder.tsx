import { useNavigate, useParams } from 'react-router-dom';
import '../styles/pages/TrackOrder.scss'
import { useEffect, useState } from 'react';
import { getOrder } from '../api';
import { Order } from '../types';

const TrackOrder = () => {

  const statusMap: Record<Order['status'], string> = {
    'new': 'Обрабатывается',
    'processing': 'Собирается',
    'shipped': 'Доставляется',
    'delivered': 'Доставлен',
  }
  const navigate = useNavigate();
  const {id} = useParams<{id: string}>();
  console.log(id)
  const [order, setOrder] = useState<Order | null>(null);
  const [loading, setLoad]  = useState(true);
  const [error, setError]   = useState<string | null>(null);

  useEffect(() => {
    if (id) {
      getOrder(Number(id)).then(setOrder).catch(error => setError(error.message)).finally(() => setLoad(false));
    }
  }, [id]);

  if (loading) return <p>Загрузка...</p>
  if (error) return <p>{error}</p>
  if (!order) return null
  return (
    <section className='track-order'>
      <p>Спасибо за ваш заказ!</p>
      <p>Номер заказа: <strong>{order.id}</strong></p>
      <p>Статус заказа: <strong>{statusMap[order.status]}</strong></p>
      <button onClick={() => navigate('/')}>Вернуться в магазин</button>
    </section>
  )
}

export default TrackOrder
