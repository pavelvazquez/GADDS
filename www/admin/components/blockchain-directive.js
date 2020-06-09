angular.module('altrs.directive.blockchain', ['altrs.service.socket', 'altrs.service.channel'])

    .directive('blockchainLog', function ($document, SocketService, ChannelService, $log) {
      return {
        restrict: 'E',
        replace: false,
        template: '<div class="bc-wrapper" id="footerWrap" ng-init="ctl.init()">' +
            '<div id="bc-wrapper-block" ng-class="ctl.getStatusClass()">' +
            '<i class="material-icons" title="{{ctl.getStatusText()}}">device_hub</i>' +

            '<div id="details" ng-show="!!ctl.blockInfo" >' +
            '<p> Block:    {{ctl.blockInfo.header.data_hash|limitTo:25}}...' +
            '<br> TXID:     {{ctl.blockInfo.data.data[0].payload.header.channel_header.tx_id|limitTo:25}}...' +
            '<br> Type:     {{ctl.blockInfo.data.data[0].payload.header.channel_header.typeString}}' +
            '<br> Created:  {{ctl.blockInfo.data.data[0].payload.header.channel_header.timestamp}}' +
            '<br> Height:   {{ctl.blockInfo.header.number}}' +
            '</p>' +
            '<hr class="line">' +
            //'<certificate title="false" data="ctl.blockInfo.data.data[0].payload.data.actions[0].header.creator.IdBytes"></certificate>' +
            '</div>' +

            '</div>' +
            '</div>',
        controllerAs: 'ctl',
        controller: function ($scope, $element) {
          var ctl = this;

          var clicked = false;
          var blockCount = 0;
          var blockWidth = 36;

          ctl.blockInfo = null;

          setInterval(function () {
            removeExtraBlocks();
          }, 2000);

          var socket = null;

          var stateClasses = {
            'error': 'red-text',
            'connected': 'light-blue-text aqua-text',
            'disconnected': 'red-text',
            'connecting': 'orange-text',
            'default': ''
          };

          /**
           *
           */
          ctl.init = function () {
            socket = SocketService.getSocket();


            socket.on('connect', function () {
              'use strict';
              //in some cases config is still loading and fetching channels will fail because of peer absence
              setTimeout(function () {
                ChannelService.list()
                    .then(function (channels) {
                      socket.emit('listen_channel', {
                        channelId: channels[0].channel_id,
                        fullBlock: true
                      });
                    });
              }, 5000);
            });

            let blocks = [];

            $log.log('chainblock event registered');
            socket.on('chainblock', function (block) {

              if (blocks.includes(block.header.number)) return;
              blocks.push(block.header.number);

              $log.log('server chainblock:', block);
              addChainblocks(block);
            });
          };

          /**
           *
           */
          ctl.getStatusClass = function () {
            return stateClasses[SocketService.getState()] || stateClasses['default'];
          };

          /**
           *
           */
          ctl.getStatusText = function () {
            return SocketService.getState();
          };

          /**
           * @param chainblock
           */
          function addChainblocks(chainblock) {
            var width = $(document).width();

            var blockElement = _blockHtml(chainblock).css({left: '+=' + width}).animate({left: '-=' + width});
            blockCount++;
            $element.find('#bc-wrapper-block').append(blockElement);
          }

          function _blockHtml(block) {
            var tx = block && block.header && block.header.data_hash;
            if (typeof (tx) == "undefined") tx = "   ";
            return $('<div class="block">' + tx.substr(0, 3) + '</div>')
                .css({left: (blockCount * blockWidth)})
                .click(_onBlockClick)
                .hover(getBlockHoverIn(block), onBlockHoverOut);
          }


          function _onBlockClick() {
            clicked = !clicked;
          }

          function getBlockHoverIn(tx) {
            // blockInfo
            return function () {
              ctl.blockInfo = tx;
              $scope.$digest();
            };
          }

          function onBlockHoverOut(/*e*/) {
            if (!clicked) {
              ctl.blockInfo = null;
              $scope.$digest();
            }
          }

          /**
           * Remove extra blocks
           */
          function removeExtraBlocks() {
            if (blockCount > 10) {
              var toRemove = blockCount - 10;
              $element.find('.block:lt(' + toRemove + ')').animate({opacity: 0}, 800, function () {
                $('.block:first').remove(); /* blocks.slice(toRemove); */
              });
              $element.find('.block').animate({left: '-=' + blockWidth * toRemove}, 800, function () {
              });
              blockCount -= toRemove;
            }
          }

        }//-controller
      };
    });

