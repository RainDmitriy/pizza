import React from 'react';
import style from './Categories.module.scss';

function Categories() {

  const [activeIndex, setActiveIndex] = React.useState(0);
  const categories = [
    'Все',
    'Мясные',
    'Вегетарианская',
    'Гриль',
    'Острые',
    'Закрытые'
  ];

  return (
    <div className={style.categories}>
      <ul>
        {categories.map((item, index) => 
          <li
            style={{marginBottom: "8px"}}
            className={activeIndex === index ? style.active : ""} 
            onClick={() => setActiveIndex(index)}
            key={item}>
            {item}
          </li>
          )
        }
      </ul>
    </div>
  )
};

export default Categories;
