import '../styles/components/ProductCard.scss'
import { Link } from "react-router-dom"

const ProductCard = ({id, img, name, price }: { id: number, img: string, name: string, price: number }) => {
  return (
    <Link to={`/product/${id}`} className="product-card__link">
      <article className="product-card">
        <img className="product-card__img" src={img} alt={`product-${id}.png`} />
        <section className="product-card__info">
          <p className="product-card__name">{name}</p>
          <p className="product-card__price">{price}&#8381;</p>
        </section>
      </article>
    </Link>
  )
}

export default ProductCard