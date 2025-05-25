export type Product = {
    id: number;
    name: string;
    price: number;
    power: string;
    temperature: string;
    type: "LED" | "incandescent" | "CFL" | "smart" | string;
    lamp_cap: string;
    image_url: string;
    description: string;
    created_at: string;
    updated_at: string;
};