import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './utils/reset.css'
import './index.scss'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
