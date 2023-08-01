import { createSlice } from "@reduxjs/toolkit";


const initialState = {
  totalPrice: 0,
  cartItems : [],
};

export const itemsSlice = createSlice({
  name: "cart",
  initialState,
  reducers: {
    updateCartItems: (state, action) => {
      state.cartItems = action.payload
    },

    updateTotalPrice: (state, action) => {
      state.totalPrice = action.payload
    }
  }
});

export const { updateCartItems, updateTotalPrice } = itemsSlice.actions;
export default itemsSlice.reducer;

