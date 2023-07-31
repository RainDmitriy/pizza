import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  sortType: 0,
  filterType: 0,
  isSortActive: false
}

export const filterSlice = createSlice({
  name: "filter",
  initialState,
  reducers: {
    changeSortType: (state, action) => {
      state.sortType = action.payload;
    },
    changeFilterType: (state, action) => {
      state.filterType = action.payload;
    },

    sortToggle: (state) => {
      state.isSortActive = !state.isSortActive;
    }
  }
})

export const { changeSortType, changeFilterType, sortToggle } = filterSlice.actions;

export default filterSlice.reducer;