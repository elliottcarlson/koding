class NewMemberBucketData extends KDObject

  constructor:(options, @buckets)->

    @bongo_ = {}
    @meta   = @buckets[0].meta
    @bongo_.constructorName = "NewMemberBucketData"
    super


class NewMemberBucketView extends JView

  constructor: (options = {}, data)->

    options.cssClass = "new-member"

    super

    @group = new LinkGroup {}, @getData().buckets.map (bucket)-> bucket.anchor

  viewAppended:->

    super

    @timer = @utils.wait 800, =>
      @$('.fx').removeClass "out hidden"
      @timer = @utils.wait 400, =>
        @$('.fx').addClass "hidden"

  pistachio:->
    """
    <span class='icon fx out'></span>
    <span class='icon'></span>
    {{> @group}}
    <span class='action'>became member.</span>
    """




# OLD

class NewMemberBucketItemView extends KDView

  constructor:(options,data)->
    options = $.extend options,
      cssClass : "new-member"
    super options,data

    @anchor = new ProfileLinkView origin: data.anchor

  render:->

  addCommentBox:->

  viewAppended:->
    @setTemplate @pistachio()
    @template.update()

  pistachio:->
    """
    <span class='icon'></span>
    {{> @anchor}}
    <span class='action'>became a member.</span>
    """
