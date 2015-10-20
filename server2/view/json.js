module.exports = function(handler){
	return async function(req,res,next){
		try{
			var result = await handler(req,res,next);
			res.send(JSON.stringify({
				code:0,
				msg:'',
				data:result
			}));
		}catch(e){
			res.send(JSON.stringify({
				code:1,
				msg:e,
				data:null
			}));
		}
	};
};