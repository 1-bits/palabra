<?php
/**
 * @Name: Consulta biblegateway
 * @author : Cesar Darinel Ortiz
 * @version: 0.0.1
 * @date : 2018-10-15
 *
 **/
class Verso{

	private $_fileName= 'config.json';
	private $_dataJson;
	private $_url;
    /**
	 * 
	 */
	public function __construct($book,$version,$capituli_verso){
		$handle = fopen($this->_fileName, 'r');
		$this->_dataJson = json_decode(fread($handle, filesize($this->_fileName)), true);
		fclose($handle);
		$this->_url = $this->_dataJson{'url'}.$this->_dataJson{'search'}
					 .$book.$capituli_verso. $this->_dataJson{'version'}.$version;
	}

	public function _getResource(){
		
		$c = curl_init($this->_url);
		curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
		curl_setopt($c, CURLOPT_PROXY, "http://acapsrv02d");
		curl_setopt($c, CURLOPT_PROXYPORT,"3128"); 
		curl_setopt($c, CURLOPT_SSL_VERIFYPEER, false);
		$html = curl_exec($c);

		if (curl_error($c))
			die(curl_error($c));
		$status = curl_getinfo($c, CURLINFO_HTTP_CODE);
		curl_close($c);

		// create new DOMDocument
		$dom = new \DOMDocument('1.0', 'UTF-8');
		// set error level
		$internalErrors = libxml_use_internal_errors(true);
		// load HTML
		$dom->loadHTML($html);
		// Restore error level
		libxml_use_internal_errors($internalErrors);
		$xpath = new DomXpath($dom);
		$tr = $xpath->query('//div[@class="passage-text"]//p//span')->item(0);
		foreach ($tr->childNodes as $node) {
			$string[] = $node->nodeValue;
		}
		$texto=implode("",$string);
	  	return preg_replace('/[0-9]+/', '', $texto);
	}
}