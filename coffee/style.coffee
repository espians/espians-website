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
            width: '900px'
