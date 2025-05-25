import '../styles/pages/Home.scss'
import { products } from '../data/products'
import ProductCard from '../components/ProductCard'

const Home = () => {
  return (
    <main className="catalog">
      <div className='catalog__grid'>
        {products.map( (product) => ( <ProductCard key = {product.id} id={product.id} img={product.image_url} name={product.name} price={product.price} /> ) )}
      </div>
    </main>
  )
}

export default Home
