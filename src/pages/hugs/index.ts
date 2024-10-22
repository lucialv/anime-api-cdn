import type { APIRoute } from "astro";

export const GET: APIRoute = ({ request }) => {
  const totalGifs = 30;
  const randomNumber = Math.floor(Math.random() * totalGifs) + 1;
  const gifUrl = `https://cdn.lucia-dev.com/hug${randomNumber
    .toString()
    .padStart(2, "0")}.gif`; // URL of the GIF

  // Return the URL of the GIF
  return new Response(JSON.stringify({ url: gifUrl }), {
    headers: {
      "Content-Type": "application/json",
    },
  });
};
