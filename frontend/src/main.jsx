import React from 'react';
import { createRoot } from 'react-dom/client'
import { Toaster } from 'react-hot-toast'

import 'react-toastify/dist/ReactToastify.css';
import './style.css';

import Router from './router';

const routes = (
  <React.StrictMode>
    <Router />
    <Toaster />
  </React.StrictMode>
)

const container = document.getElementById("root");
const root = createRoot(container);
root.render(routes)
