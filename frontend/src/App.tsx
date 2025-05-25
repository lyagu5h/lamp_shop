import { BrowserRouter as Router,Route, Routes } from 'react-router-dom'
import './App.scss'
import NotFound from './pages/NotFound'
import ProductDetails from './pages/ProductDetails'
import Cart from './pages/Cart'
import TrackOrder from './pages/TrackOrder'
import Home from './pages/Home'
import Header from './components/Header'
import Footer from './components/Footer'

function App() {

  return (
    <Router>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/product/:id" element={<ProductDetails/>} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/orders/:orderId/status" element={<TrackOrder />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
      <Footer/>
    </Router>
  )
}

export default App
