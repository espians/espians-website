# Public Domain (-) 2014 The Espians Website Authors.
# See the Espians Website UNLICENSE file for details.

module.exports = (api) ->

    api.add

        '*':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'

        body:
            background: '#f9f9f9'
            margin: 0
            padding: 0
            lineHeight: 1.3
            paddingTop: '55px'

        h1:
            fontFamily: 'Montserrat'
            fontWeight: '400'
            fontSize: '17px'
            color: '#2b2b2b'
            letterSpacing: '0.8px'

        h2:
            fontFamily: 'Montserrat'
            fontSize: '26px'
            fontWeight: 400
            lineHeight: 1.3
            color: '#2b2b2b'
            clear: 'both'
            textAlign: 'center'
            marginTop: '40px'

        h3:
            fontFamily:'Merriweather Sans'
            fontSize: '21px'
            fontWeight: 400
            lineHeight: 1.3
            marginBottom: 0

        p:
            fontFamily: 'Merriweather'
            fontSize:'14px'
            fontWeight: 300
            padding: 0
            margin: 0

        a:
            fontFamily: 'Merriweather'
            fontSize:'14px'
            fontWeight: 400
            color: '#cc0000'
            textDecoration: 'none'

        '#full-logo':
            display: 'inline-block'
            width: '150px'
            marginTop: '5px'
            marginBottom: '2px'

        '.logo':
            width: '38px'
            float: 'left'
            marginRight: '5px'
            paddingTop: '3px'
            paddingBottom: '3px'

         header:
            width: '100%'
            height: '55px'
            backgroundColor: '#f9f9f9'
            top: 0
            zIndex: 3000
            transition: 'top 0.2s ease-in-out'
            position: 'fixed'
            boxShadow: '0 1px 10px rgba(0, 0, 0, 0.1)'

        '.nav-up':
            top: '-55px'

        '#nav':
            float: 'right'

        '#nav>ul>li>a':
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '16px'
            paddingLeft: '10px'
            paddingRight: '14px'
            color: '#2b2b2b'
            cursor: 'pointer'

        '#network':
            backgroundColor: '#878787'
            color: '#ffffff'
            height: '525px'
            overflow: 'hidden'
            position: 'relative'
            paddingBottom: 0
            paddingTop: 0

        '#nodes':
            backgroundColor: 'transparent'
            position: 'absolute'
            top: 0
            width: '1400px'

            rect:
                fill: 'none'
                pointerEvents: 'all'

            '.cursor':
                fill: 'none'
                pointerEvents: 'none'
                stroke: 'brown'

            '.link':
                stroke: 'rgba(204,0,0,0.6)'
                strokeWidth: 1
                transitionDuration: '1s'
                transitionProperty: 'stroke fill'
                zIndex: 0

            '.link.active':
                stroke: 'rgba(204,0,0,0.9)'
                transitionDuration: '1s'
                transitionProperty: 'stroke'

            '.node':
                fill: '#878787'
                stroke: '#bbb'
                transitionDuration: '1s'
                transitionProperty: 'stroke fill'
                zIndex: 100

            '.node.active':
                stroke: '#dd4e58'
                transitionDuration: '1s'
                transitionProperty: 'stroke fill'

        '.wrapper':
            margin: '0 auto'
            width: '990px'
            clear: 'left'

        '#tagline':
            fontFamily: 'Merriweather Sans'
            fontSize: '32px'
            fontWeight: 400
            lineHeight: 1.4
            textShadow: '1px 1px 3px #2b2b2b'
            marginTop: '225px'
            position: 'relative'
            zIndex: 20

        '#calltoaction':
            cursor: 'pointer'
            fontFamily: 'Merriweather Sans'
            fontSize: '26px'
            color: '#2b2b2b'
            fontWeight: 400
            lineHeight: 1.4
            backgroundColor: 'rgba(250,250,250,0.6)'
            position: 'relative'
            zIndex: 20
            border: '1px solid #bbb'
            padding: '5px 10px'
            textAlign: 'center'
            marginTop: '40px'
            marginLeft: 'auto'
            marginRight: 'auto'
            width: '180px'
            height: '50px'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'

        '#calltoaction:hover':
            border: '1px solid #cc0000'

        '.card':
            position: 'relative'
            height: '450px'
            width: '300px'
            backgroundColor: '#f4f4f4'
            borderBottom: '2px solid #e5e5e6'
            boxShadow: '0 -2px 10px rgba(0,0,0,0.1)'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'
            marginBottom: '40px'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'

         '.card-centered':
            float: 'none !important'
            margin: '0 auto !important'

         '.card:hover':
            borderBottom: '2px solid #cc0000'

         '.card-img':
            width: '300px'
            height: '190px'
            backgroundColor: 'rgba(255,255,255,0.8)'

         '.card-text':
            marginLeft: '20px'
            marginRight: '20px'

         '.card-text>h3':
            marginBottom: '8px'
            marginTop: '14px'

         '.card-email':
            marginTop: '12px'

         img:
            width: '100%'
            height: '100%'

         ul:
           listStyle: 'none'

         li:
           display: 'inline-block'
           clear: 'left'

         '.card-smedia':
           position: 'absolute'
           bottom: 0
           right: 0
           width: '85px'
           height: '30px'
           paddingTop: 0
           marginTop: 0
           zIndex: 2

         '.icon':
            width: '24px'
            paddingLeft: '4px'
            marginRight: '6px'
            marginLeft: '-6px'
            float: 'right'
            opacity: 0.5

          '.icon:hover':
            opacity: 1

          '.icon:hover + .card-followus':
            color: 'rgba(0,0,0,0.5)'

          '.person-card':
            position: 'relative'
            height: '450px'
            width: '300px'
            marginBottom: '20px'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'

          '.person-card-centered':
             float: 'none !important'
             margin: '0 auto !important'

          '.avatar':
            borderRadius: '50%'
            width: '190px'
            height: '190px'
            marginLeft: '50px'
            marginRight: '50px'

          '.person-card>*':
            textAlign: 'center'

          '.person-smedia':
            width: '112px'
            marginLeft: '98px'
            marginRight: '90px'
            height: '30px'
            marginTop: '12px'

          '.card-timeline':
            position: 'relative'
            height: '550px'
            width: '350px'
            backgroundColor: '#f4f4f4'
            borderBottom: '2px solid #e5e5e6'
            boxShadow: '0 -1px 10px rgba(0,0,0,0.1)'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'

          '.timeline-img':
             width: '350px'
             height: '190px'
             backgroundColor: 'rgba(255,255,255,0.8)'

          '.card-timeline:hover':
            borderBottom: '2px solid #cc0000'

          '.timeline-year':
            width: '75px'
            height: '32px'
            backgroundColor: '##F7F7F7'
            marginLeft: 'auto'
            marginRight: 'auto'
            borderBottom: '2px solid rgba(204,0,0,1)'
            clear: 'both'

          '.timeline-year>p':
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '16px'
            color: '#2b2b2b'
            textAlign: 'center'
            letterSpacing: '1.2px'
            paddingTop: '8px'
            clear: 'both'

          '#timeline-divider':
            position: 'absolute'
            width: '4px'
            height: '590px'
            backgroundColor: '#939393'
            marginLeft: '493px'
            marginRight: '493px'
            zIndex: 1

          '.card-left':
            float: 'left'
            marginLeft: '103px'
            paddingTop: '15px'
            paddingBottom: '25px'
            zIndex: 2

          '.card-right':
            float: 'right'
            marginRight: '103px'
            paddingTop: '15px'
            paddingBottom: '25px'
            zIndex: 2
