# Public Domain (-) 2014 The Espians Website Authors.
# See the Espians Website UNLICENSE file for details.

module.exports = (api) ->

    api.add

        # '*':
        #     WebkitBoxSizing: 'border-box'
        #     MozBoxSizing: 'border-box'
        #     boxSizing: 'border-box'

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
            textAlign: 'center'
            marginTop: '50px'
            marginBottom: '50px'
            clear: 'both'

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

        'a::selection':
           color: '#fff'

        '::selection':
            background: 'rgba(204,0,0,0.7)'

        tbody:
            # borderSpacing: '0px !important'
            # borderCollapse: 'collapse'
            width: '100%'

        tr:
            width: '100%'

        td:
            paddingTop: '0px'
            marginTop: '0px'

        footer:
            clear:'both'

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
            transition: 'top 0.4s ease-in-out'
            position: 'fixed'
            boxShadow: '0 1px 10px rgba(0, 0, 0, 0.1)'

        '.nav-up':
            top: '-55px'

        '#nav':
            float: 'right'

        '#nav>ul>li':
            marginTop: '-3px'

        '#nav>ul>li>a':
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '16px'
            paddingLeft: '10px'
            paddingRight: '14px'
            color: '#2b2b2b'
            cursor: 'pointer'

        '#nav>ul>li>a:hover':
            color: '#cc0000'

        '.header-smedia':
            borderLeft: '2px solid #cc0000'
            zIndex: 3000
            top: 0
            clear: 'all'
            marginTop: '-2px'

        '.icon-header':
            width: '18px'
            marginLeft: '12px'
            marginTop: '5px'
            opacity: 0.6

        '.icon-header:hover':
            opacity: 1

        '#icon-header-fb':
            marginLeft: '4px'

        '#network':
            backgroundColor: '#878787'
            color: '#ffffff'
            height: '460px'
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
            marginTop: '192px'
            position: 'relative'
            zIndex: 20

        '#calltoaction':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            cursor: 'pointer'
            backgroundColor: 'rgba(250,250,250,0.6)'
            position: 'relative'
            zIndex: 20
            border: '1px solid #bbb'
            padding: '4px 10px'
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

        '#calltoaction>a':
            fontFamily: 'Merriweather Sans'
            fontSize: '26px'
            color: '#2b2b2b'
            fontWeight: 400
            lineHeight: 1.4

        '.card':
            position: 'relative'
            height: '510px'
            width: '300px'
            backgroundColor: '#f4f4f4'
            borderBottom: '2px solid #e5e5e6'
            boxShadow: '0 -2px 10px rgba(0,0,0,0.1)'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'
            marginBottom: '50px'

        '.card-centered':
            float: 'none !important'
            margin: '0 auto !important'

        '.center-wrapper':
            width:'660px'
            height:'auto !important'
            float: 'center !important'
            marginLeft: 'auto !important'
            marginRight: 'auto !important'

        '.card:hover':
            borderBottom: '2px solid #cc0000'

        '.card-img':
            width: '300px'
            height: '190px'
            backgroundColor: 'rgba(255,255,255,0.8)'

        '.card-text':
            lineHeight: '1.4em'
            paddingBottom: '50px'
            paddingLeft: '20px'
            paddingRight: '20px'

        '.clearfix:before, .clearfix:after':
           content: " "
           display: 'table'

        '.clearfix:after':
           clear: 'both'

        '.person-text':
            paddingBottom: '10px'
            paddingLeft: '20px'
            paddingRight: '20px'
            lineHeight: '1.3em'

        '.card-text>h3':
            lineHeight: '1em'
            paddingBottom: '18px'
            paddingTop: '24px'
            marginTop: '0px'

        '.person-text>p':
            paddingBottom: '8px'
            paddingTop: '14px'

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
           WebkitBoxSizing: 'border-box'
           MozBoxSizing: 'border-box'
           boxSizing: 'border-box'
           position: 'absolute'
           bottom: 0
           right: 0
           marginRight: '3px'
           width: '120px'
           height: '30px'
           paddingTop: 0
           marginTop: 0
           zIndex: 2

        '.icon':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            width: '23px'
            paddingLeft: '4px'
            marginRight: '6px'
            marginLeft: '-6px'
            float: 'right'
            opacity: 0.5

          '.icon:hover':
            opacity: 1

          '#clients':
            marginTop: '50px'
            clear: 'both'

        #   '.buffer':
        #     marginTop: '589px'
        #     height: '1px'
        #     width: '990px'
        #     marginLeft: 'auto'
        #     marginRight: 'auto'
        #     position: 'relative'
        #     backgroundColor: '#f9f9f9'

          '.clients':
            width: '990px'
            height: '331px'

          '.person-card':
            width: '300px'
            marginBottom: '20px'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'

          '.avatar':
            borderRadius: '50%'
            width: '190px'
            height: '190px'
            marginLeft: '50px'
            marginRight: '50px'

          '.avatar-advisor':
            borderRadius: '50%'
            height: '120px'
            width: '120px'
            marginLeft: '85px'
            marginRight: '85px'

          '.person-card>*':
            textAlign: 'center'

          '.person-smedia':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            clear:'both'
            width: '120px'
            marginLeft: 'auto'
            marginRight: 'auto'
            height: '30px'
            marginBottom:'10px'

          '.icon-person':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            width: '23px'
            paddingLeft: '4px'
            marginRight: '6px'
            marginLeft: '-6px'
            display: 'inline-block'
            opacity: 0.5

          '.icon-person:hover':
            opacity: 1

          '.timeline-sections':
            height: '622px'

          '.card-timeline':
            position: 'relative'
            width: '350px'
            backgroundColor: '#f4f4f4'
            borderBottom: '2px solid #e5e5e6'
            boxShadow: '0 -1px 10px rgba(0,0,0,0.1)'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'

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

          '.timeline-divider':
            width: '4px'
            backgroundColor: '#939393'

          '.timeline-img':
            width: '350px'
            height: '190px'
            backgroundColor: 'rgba(255,255,255,0.8)'

          'table.timeline-wrapper':
            width: '100%'

          '.card-left':
            paddingLeft: '103px'
            paddingTop: '15px'
            paddingBottom: '25px'
            width: '390px'

          '.card-right':
            paddingLeft: '40px'
            paddingTop: '15px'
            marginBottom: '25px'
            width: '453px'

          '.footer-wrapper':
            paddingTop: '10px'
            height: '600px'
            paddingBottom:0
            backgroundColor: '#f4f4f4'

          '.contact-wrapper':
            width: 'auto'
            marginLeft: 'auto'
            marginRight:'auto'

          '.contact-text':
            float: 'left'
            marginLeft:'22%'

          '.contact-text>h3':
            marginTop: 0
            marginBottom:'18px'

          '.contact-text>p':
            lineHeight: 1.4

          '.contact-smedia':
            marginTop: '20px'
            float: 'left'
            marginLeft: '-12px'

          '.map':
            width:"50%"
            float: 'right'
