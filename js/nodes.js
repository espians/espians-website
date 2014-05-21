// Adapted from the wonderful animation used on the Serf Website.
//
// Original source:
// https://github.com/hashicorp/serf/blob/master/website/source/javascripts/app/nodes.js
//
// Licensed under the Mozilla Public License:
// https://github.com/hashicorp/serf/blob/master/LICENSE

(function (root) {

    var width = 1400,
        height = 490,
        border = 50,
        numberNodes = 128,
        linkGroup = 0;

    var nodes = [];
    for (i=0; i<numberNodes; i++) {
        nodes.push({
            x: Math.random() * (width - border) + (border / 2),
            y: Math.random() * (height - border) + (border / 2),
        });
    }

    var fill = d3.scale.category20();

    var force = d3.layout.force()
        .size([width, height])
        .nodes(nodes)
        .linkDistance(60)
        .charge(-1)
        .gravity(0.0004)
        .on("tick", tick);

    var svg = d3.select("#network").append("svg")
        .attr('id', 'nodes')
        .attr("width", width)
        .attr("height", height)

    // set left value after adding to the dom
    resize();

    svg.append("rect")
        .attr("width", width)
        .attr("height", height);

    var nodes = force.nodes(),
        links = force.links(),
        node = svg.selectAll(".node"),
        link = svg.selectAll(".link");

    var cursor = svg.append("circle")
        .attr("r", 30)
        .attr("transform", "translate(-100,-100)")
        .attr("class", "cursor");

    function createLink(index) {
        var node = nodes[index];
        var nodeSelected = svg.select("#id_" + node.index).classed("active linkgroup_"+ linkGroup, true);
        var distMap = {};
        var distances = [];
        for (var i=0; i<nodes.length; i++) {
            if (i == index) {
                continue
            }
            var target = nodes[i];
            var selected = svg.select("#id_" + i);
            var dx = selected.attr('cx') - nodeSelected.attr('cx');
            var dy = selected.attr('cy') - nodeSelected.attr('cy');
            var dist = Math.sqrt(dx * dx + dy * dy)
            if (dist in distMap) {
                distMap[dist].push(target)
            } else {
                distMap[dist] = [target]
            }
            distances.push(dist)
        }
        distances.sort(d3.ascending);
        var n = Math.floor(Math.random() * (5 - 2 + 1) + 2);
        // var n = 3;
        for (i = 0; i < n; i++) {
            var dist = distances[i]
            var target = distMap[dist].pop()
            var link  = {
                source: node,
                target: target
            }
            links.push(link);
        }
        restart();
    }

    function tick() {
        link.attr("x1", function(d) { return d.source.x; })
            .attr("y1", function(d) { return d.source.y; })
            .attr("x2", function(d) { return d.target.x; })
            .attr("y2", function(d) { return d.target.y; });
        node.attr("cx", function(d) { return d.x; })
            .attr("cy", function(d) { return d.y; });
    }

    function restart() {
        node = node.data(nodes);
        node.enter().insert("circle", ".cursor")
            .attr("class", "node")
            .attr("r", 9)
            .attr("id", function (d, i) {
                return ("id_" + i)
            })
            .call(force.drag);
        link = link.data(links);
        link.enter().insert("line", ".node")
            .attr("class", "link active linkgroup_"+ linkGroup);
        force.start();
        resetLink(linkGroup);
        linkGroup++;
    }

    function resetLink(num){
        setTimeout(resetColors, 700, num)
    }

    function resetColors(num){
        svg.selectAll(".linkgroup_"+ num).classed('active', false)
    }

    root.onresize = resize;

    function resize() {
        var nodeC = document.getElementById('nodes');
            wW = root.innerWidth;
        nodeC.style.left = ((wW - width) / 2 ) + 'px';
    }

    function init() {
        restart();
        for (i=0;i<numberNodes;i++) {
            setTimeout(createLink, 20*i+10, i);
            // setTimeout(createLink, 700*i+1000, i);
        }
    }

    init();

})(window);
