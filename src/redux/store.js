import { configureStore, combineReducers } from '@reduxjs/toolkit';
import filter from './slices/filterSlice';
import items from './slices/itemsSlice';
import cart from './slices/cartSlice';

const redusers = combineReducers({
  filter,
  items,
  cart,
});

export const store = configureStore({
  reducer: redusers,
});
