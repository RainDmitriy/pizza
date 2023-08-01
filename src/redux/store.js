import { configureStore, combineReducers } from "@reduxjs/toolkit";
import filterReducer from "./slices/filterSlice";
import itemsReducer from "./slices/itemsSlice";
import cartReducer from "./slices/cartSlice";

const redusers = combineReducers({
  filter: filterReducer,
  items: itemsReducer,
  cart: cartReducer
})

export const store = configureStore({
  reducer: redusers
});