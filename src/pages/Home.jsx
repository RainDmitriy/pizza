import React from 'react';
import Header from '../components/Header';
import PizzaBlock from '../components/PizzaBlock';
import Sort from '../components/Sort';
import Categories from '../components/Categories';
import ContentLoader from 'react-content-loader';
import { useSelector } from 'react-redux';

function Home() {
  const { sortType, filterType } = useSelector((state) => state.filter);
  const { items, isLoaded } = useSelector((state) => state.items);

  const sorted = (pizzas, sortType) => {
    switch (sortType) {
      case 0:
        return pizzas.toSorted((a, b) => b.Rating - a.Rating);
      case 1:
        return pizzas.toSorted((a, b) => a.Prices[0] - b.Prices[0]);
      case 2:
        return pizzas.toSorted((a, b) => a.Title.localeCompare(b.Title));
      default:
        return pizzas;
    }
  };

  const filtered = (pizzas, filterType) => {
    switch (filterType) {
      case 0:
        return pizzas;
      case 1:
        return pizzas.filter((item) => item.Props[0] === true);
      case 2:
        return pizzas.filter((item) => item.Props[1] === true);
      case 3:
        return pizzas.filter((item) => item.Props[2] === true);
      case 4:
        return pizzas.filter((item) => item.Props[3] === true);
      default:
        return pizzas;
    }
  };

  return (
    <div className="wrapper">
      <Header />
      <div className="content">
        <div className="container">
          <div className="content__top">
            <Categories />
            <Sort />
          </div>
          <h2 className="content__title">Все пиццы</h2>
          <div className="content__items">
            {isLoaded
              ? sorted(filtered(items, filterType), sortType).map((item) => (
                  <PizzaBlock key={item.ItemId} {...item} />
                ))
              : [...Array(16)].map((obj) => (
                  <ContentLoader
                    speed={2}
                    width={280}
                    height={466}
                    viewBox="0 0 280 466"
                    backgroundColor="#f3f3f3"
                    foregroundColor="#ecebeb">
                    <circle cx="130" cy="130" r="120" />
                    <rect x="0" y="260" rx="0" ry="0" width="280" height="26" />
                    <rect x="129" y="416" rx="20" ry="20" width="151" height="46" />
                    <rect x="2" y="306" rx="10" ry="10" width="276" height="90" />
                    <rect x="35" y="426" rx="0" ry="0" width="60" height="27" />
                  </ContentLoader>
                ))}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
