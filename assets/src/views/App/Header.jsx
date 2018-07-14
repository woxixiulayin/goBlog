import React, { Component } from 'react'
import { navList } from 'src/routes/config'
import { NavLink } from 'react-router-dom'

class Header extends Component {
    render() {
        return (
            <header
             className="text-center relative"
            >
                <h2>Jackson Liu's house</h2>
                <nav className="nav clearfix">
                    <div className="links">
                    {navList.map(item => 
                        <NavLink
                          exact
                          key={item.label}
                          to={item.path}
                          activeClassName="current"
                        >
                            {item.zH}
                        </NavLink>
                    )}
                    <a href="https://github.com/woxixiulayin" target="_blank">github</a>
                    </div>
                </nav>
            </header>
        )
    }
}

export default Header
