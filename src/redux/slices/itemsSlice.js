import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  isLoaded: false,
  items: [],
};

export const itemsSlice = createSlice({
  name: 'items',
  initialState,
  reducers: {
    updateItems: (state, action) => {
      state.items = action.payload;
    },

    loadToggle: (state, action) => {
      state.isLoaded = action.payload;
    },
  },
});

export const { updateItems, loadToggle } = itemsSlice.actions;
export default itemsSlice.reducer;
