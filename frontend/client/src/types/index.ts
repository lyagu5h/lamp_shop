export type Product = {
    id: number;
    name: string;
    price: number;
    power: string;
    temperature: string;
    type: "LED" | "incandescent" | "CFL" | "smart" | string;
    lamp_cap: string;
    image_url: string;
    description: {
      String: string;
      Valid: boolean
    };
    created_at: string;
    updated_at: string;
};

export type OrderItem = {
  product_id: number;
  quantity: number;
  price: number;
};

export type Order = {
  id: number;
  customer_name: string;
  customer_email: string;
  customer_phone: string;
  address: string;
  total_amount: number;
  status: string;
  items: OrderItem[];
  created_at: string;
  updated_at: string;
};