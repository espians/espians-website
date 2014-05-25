# Public Domain (-) 2014 The Espians Website Authors.
# See the Espians Website UNLICENSE file for details.

module.exports = (api) ->

    api.add

        '*':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'

        body:
            background: '#fff'
            margin: 0
            padding: 0
            lineHeight: 1.4

        h2:
            fontFamily: 'Merriweather'
            fontSize: '24px'
            fontWeight: 700
            lineHeight: 1.4
            clear: 'left'

        h3:
            fontFamily:'Merriweather Sans'
            fontSize: '21px'
            fontWeight: 400
            lineHeight: 1.4
            marginBottom: 0

        p:
            fontFamily: 'Merriweather'
            fontSize:'14px'
            fontWeight: 300
            padding: 0
            margin: 0

        a:
            fontFamily: 'Merriweather'
            fontSize:'12px'
            fontWeight: 300
            color: '#cc0000'
            textDecoration: 'none'

        '#network':
            backgroundColor: '#939393'
            color: '#ffffff'
            height: '490px'
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
                stroke: '#bbb'
                strokeWidth: 1.5
                transitionDuration: '1s'
                transitionProperty: 'stroke fill'
                zIndex: 0

            '.link.active':
                stroke: '#dd4e58'
                transitionDuration: '1s'
                transitionProperty: 'stroke'

            '.node':
                fill: '#939393'
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

        '.card':
            position: 'relative'
            height: '450px'
            width: '300px'
            backgroundColor: '#f2f2f2'
            borderBottom: '2px solid #e5e5e6'
            WebkitTransitionDuration: '500ms'
            WebkitTransitionProperty: 'border'
            WebkitTransitionTimingFunction: 'ease-in-out'
            marginBottom: '20px'
            display: 'inline-block'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'

         '.card:hover':
            borderBottom: '2px solid #cc0000'

         '.card-img':
            width: '300px'
            height: '190px'
            backgroundColor: '#e5e5e6'

         '.card-text':
            marginLeft: '20px'
            marginRight: '20px'

         '.card-text>h3':
            marginBottom: '12px'

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

         '.icon':
            width: '24px'
            paddingRight: '4px'
            float: 'Left'
            opacity: 0.5

          '.icon:hover':
            opacity: 1

          '.person-card':
            position: 'relative'
            height: '450px'
            width: '300px'
            marginBottom: '20px'
            display: 'inline-block'
            float: 'left'
            marginLeft: '15px'
            marginRight: '15px'

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
