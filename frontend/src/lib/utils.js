import { clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs) {
  return twMerge(clsx(inputs))
}

export function RandomColors() {
  const numColors = 10;
  const colors = [];

  for (let i = 0; i < numColors; i++) {
    const randomRed = Math.floor(Math.random() * 256);
    const randomGreen = Math.floor(Math.random() * 256);
    const randomBlue = Math.floor(Math.random() * 256);
    const randomColor = `rgba(${randomRed}, ${randomGreen}, ${randomBlue}, 0.5)`;
    colors.push(randomColor);
  }

  return colors;
}