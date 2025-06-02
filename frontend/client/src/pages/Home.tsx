import '../styles/pages/Home.scss'
import { useEffect, useState } from 'react'
import { getProducts } from '../api'
import ProductCard from '../components/ProductCard'
import { Product } from '../types'

const Home = () => {
  const [products, setProducts] = useState<Product[]>([])

  useEffect(() => {
    getProducts().then(setProducts)
  }, [])

  return (
    <main className="catalog">
      <div className='catalog__grid'>
        {products.map( (product) => ( <ProductCard key = {product.id} id={product.id} img={product.image_url} name={product.name} price={product.price} /> ) )}
      </div>
    </main>
  )
}

export default Home
