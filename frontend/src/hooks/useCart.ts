import { useEffect, useState } from 'react';

export interface CartItem {
  id: number;
  name: string;
  price: number;
  quantity: number;
  image_url: string;
}

export const useCart = () => {
  const [cart, setCart] = useState<CartItem[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadCart = () => {
      try {
        const storedCart = localStorage.getItem('cart');
        if (storedCart) {
          setCart(JSON.parse(storedCart));
        }
      } catch (e) {
        if (e instanceof Error) {
            setError(e.message);
          } 
      }
    };
    loadCart();
  }, []);

  const updateCart = (newCart: CartItem[]) => {
    setCart(newCart);
    localStorage.setItem('cart', JSON.stringify(newCart));
  };

  const addToCart = (product: Omit<CartItem, 'quantity'>, quantity = 1) => {
    const existingItem = cart.find(item => item.id === product.id);
    
    const updatedCart = existingItem
      ? cart.map(item =>
          item.id === product.id
            ? { ...item, quantity: item.quantity + quantity }
            : item
        )
      : [...cart, { ...product, quantity }];
    
    updateCart(updatedCart);
  };

  const removeFromCart = (id: number) => {
    updateCart(cart.filter(item => item.id !== id));
  };

  const updateQuantity = (id: number, quantity: number) => {
    if (quantity < 1) return;
    updateCart(
      cart.map(item => (item.id === id ? { ...item, quantity } : item))
    );
  };

  const calculateTotal = () => {
    return cart.reduce((total, item) => total + item.price * item.quantity, 0);
  };

  return {
    cart,
    error,
    addToCart,
    removeFromCart,
    updateQuantity,
    calculateTotal,
  };
};