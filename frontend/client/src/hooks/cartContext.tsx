import {
  createContext,
  useEffect,
  useReducer,
  type ReactNode,
} from 'react'

export interface CartItem {
  id: number
  name: string
  price: number
  image_url: string
  quantity: number
}

type Action =
  | { type: 'ADD'; item: Omit<CartItem, 'quantity'>; qty?: number }
  | { type: 'REMOVE'; id: number }
  | { type: 'SET_QTY'; id: number; qty: number }
  | { type: 'CLEAR' }

function reducer(state: CartItem[], act: Action): CartItem[] {
  switch (act.type) {
    case 'ADD': {
      const exists = state.find(i => i.id === act.item.id)
      return exists
        ? state.map(i =>
            i.id === act.item.id
              ? { ...i, quantity: i.quantity + (act.qty ?? 1) }
              : i
          )
        : [...state, { ...act.item, quantity: act.qty ?? 1 }]
    }
    case 'REMOVE':
      return state.filter(i => i.id !== act.id)
    case 'SET_QTY':
      return state.map(i =>
        i.id === act.id ? { ...i, quantity: act.qty } : i
      )
    case 'CLEAR':
      return []
    default:
      return state
  }
}

export interface CartCtx {
  cart: CartItem[]
  clear: () => void
  add: (item: Omit<CartItem, 'quantity'>, qty?: number) => void
  remove: (id: number) => void
  setQty: (id: number, qty: number) => void
  total: () => number
}

const CartContext = createContext<CartCtx | null>(null)

export const CartProvider = ({ children }: { children: ReactNode }) => {
  const [cart, dispatch] = useReducer(
    reducer,
    [],
    () => JSON.parse(localStorage.getItem('cart') ?? '[]')
  )

  useEffect(() => {
    localStorage.setItem('cart', JSON.stringify(cart))
  }, [cart])

  const value: CartCtx = {
    cart,
    clear: () => dispatch({ type: 'CLEAR' }),
    add: (item, qty) => dispatch({ type: 'ADD', item, qty }),
    remove: id => dispatch({ type: 'REMOVE', id }),
    setQty: (id, qty) => dispatch({ type: 'SET_QTY', id, qty }),
    total: () => cart.reduce((s, i) => s + i.price * i.quantity, 0),
  }

  return <CartContext.Provider value={value}>{children}</CartContext.Provider>
}

export { CartContext }