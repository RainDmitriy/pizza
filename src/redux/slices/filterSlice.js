import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  sortType: 0,
  filterType: 0,
};

export const filterSlice = createSlice({
  name: 'filter',
  initialState,
  reducers: {
    changeSortType: (state, action) => {
      state.sortType = action.payload;
    },
    changeFilterType: (state, action) => {
      state.filterType = action.payload;
    },
  },
});

export const { changeSortType, changeFilterType } = filterSlice.actions;

export default filterSlice.reducer;
