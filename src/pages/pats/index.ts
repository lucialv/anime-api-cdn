import type { APIRoute } from 'astro';

export const GET: APIRoute = ({ request }) => {
  const totalGifs = 30; // Total de GIFs disponibles
  const randomNumber = Math.floor(Math.random() * totalGifs) + 1; // Número aleatorio entre 1 y 30
  const gifUrl = `https://cdn.lucia-dev.com/pat${randomNumber.toString().padStart(2, '0')}.gif`; // URL del GIF

  return {
    body: {
      url: gifUrl, // Devuelve la URL en un objeto JSON
    },
  };
};
