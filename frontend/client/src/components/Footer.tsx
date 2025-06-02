import '../styles/components/Footer.scss'

const Footer = () => {
  return (
    <footer className="footer">
      <h1 className="footer__logo">Lamp</h1>
      <section className='footer__contacts'>
        <a className='footer__tel' href="tel: 8-800-224-2957">8-800-224-2957</a>
        <li className='footer__social-medias'>
          <ul className="footer__link"><a href="/" target="_blank">VK</a></ul>
          <ul className="footer__link"><a href="/" target="_blank">X</a></ul>
          <ul className="footer__link"><a href="/" target="_blank">TG</a></ul>
        </li>
      </section>
    </footer>
  )
}

export default Footer